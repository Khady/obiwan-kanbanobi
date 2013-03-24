package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"github.com/dchest/uniuri"
	"net"
	"time"
)

type Session struct {
	Id          uint32
	User_id     string
	Ident_date  time.Time
	Session_key string
}

// Cette fonction doit gerer l'identification d'un utilisateur (verifier qu'il n'est pas deja identifie,
// que son nom et son mdp sont valides...)
// Pour le moment, un message bateau est envoye pour dire que tout s'est bien passe
// Il faudrait que cette fonction mette aussi a jour un tableau avec des duo user/connexion pour les moments ou il
// faudra envoyer des messages a tout le monde
func MsgIdentConnect(conn net.Conn, msg *message.Msg) {
	sessionId := uniuri.New()
	u := User{Name: *msg.Ident.Login}
	var err error
	var checkPassword bool
	if err = u.GetByName(dbPool); err == nil {
		session := &Session{
			Id:          0,
			User_id:     *msg.Ident.Login,
			Ident_date:  time.Now(),
			Session_key: sessionId,
		}
		checkPassword, err = u.CheckPassword(dbPool, *msg.Ident.Password)
		if err == nil && checkPassword == true {
			err = session.Add(dbPool)
		}
	}
	var answer *message.Msg
	if err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_IDENT.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(0),
			SessionId: proto.String(*msg.Ident.Login),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(51), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_IDENT.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(u.Id),
			SessionId: proto.String(sessionId),
			Ident: &message.Msg_Ident{
				Login: proto.String(*msg.Ident.Login),
			},
		}
		CONNECTION_LIST.add(u.Id, *msg.Ident.Login, conn)
	}
	sendKanbanMsg(conn, answer)
}

func MsgIdentDisconnect(conn net.Conn, msg *message.Msg) {
	u := User{Name: *msg.Ident.Login}
	var err error
	if err = u.GetByName(dbPool); err == nil {
		session := &Session{
			User_id: *msg.Ident.Login,
		}
		err = session.DelByUserName(dbPool)
	}
	var answer *message.Msg
	if err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_IDENT.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(0),
			SessionId: proto.String(*msg.Ident.Login),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(52), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_IDENT.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(u.Id),
			SessionId: proto.String(*msg.SessionId),
			Ident: &message.Msg_Ident{
				Login: proto.String(*msg.Ident.Login),
			},
		}
		CONNECTION_LIST.del(u.Id)
	}
	sendKanbanMsg(conn, answer)
	// CONNECTION_LIST.del(*msg.AuthorId)
}

func MsgIdentIsUnidentified(conn net.Conn, msg *message.Msg) bool {
	s := Session{Session_key: *msg.SessionId}
	if err := s.GetUserSessionBySessionId(dbPool); err != nil {
		return false
	}
	return true
}

// Cette fonction a une gestion synchrone des messages (traitement les uns apres les autres, pas de traitements paralleles)
// Il faut faire une pool de worker, un dispacher et lancer l'operation a effectuer dans le dispatch.
func MsgIdent(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CONNECT:
		MsgIdentConnect(conn, msg)
	case message.CMD_DISCONNECT:
		MsgIdentDisconnect(conn, msg)
	}
}
