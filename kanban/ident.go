package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"net"
	"time"
)

type Session struct {
	Id          uint32
	User_id     string
	ident_date  time.Time
	session_key string
}

// Cette fonction doit gerer l'identification d'un utilisateur (verifier qu'il n'est pas deja identifie,
// que son nom et son mdp sont valides...)
// Pour le moment, un message bateau est envoye pour dire que tout s'est bien passe
// Il faudrait que cette fonction mette aussi a jour un tableau avec des duo user/connexion pour les moments ou il
// faudra envoyer des messages a tout le monde
func MsgIdent(conn net.Conn, msg *message.Msg) {
	test := &message.Msg{
		Target:    message.TARGET_IDENT.Enum(),
		Command:   message.CMD_SUCCES.Enum(),
		AuthorId:  proto.Uint32(1),
		SessionId: proto.String("superchainedesession"),
		Ident: &message.Msg_Ident{
			Login: proto.String(*msg.Ident.Login),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	LOGGER.Print("MsgIdent", len(data), data)
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}
