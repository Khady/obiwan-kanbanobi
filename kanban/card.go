package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

type Card struct {
	Id         int
	Name       string
	Content    string
	Column_id  uint32
	Project_id uint32
	Tags       []string
	User_id    uint32
	Scripts_id []uint32
	Write      []uint32
}

// msg.Cards.UserId est utilise par defaut pour le moment. Mais c'est un champ optionnel.
// Il faudrait faire un test pour savoir si c'est le author_id ou lui qui est utilise.
func MsgCardCreate(conn net.Conn, msg *message.Msg) {
	card := &Card{
		0,
		*msg.Cards.Name,
		*msg.Cards.Desc,
		*msg.Cards.ColumnId,
		*msg.Cards.ProjectId,
		msg.Cards.Tags,
		*msg.Cards.UserId,
		msg.Cards.ScriptsIds,
		msg.Cards.Write,
	}
	var answer *message.Msg
	if err := card.Add(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
		Target:    message.TARGET_CARDS.Enum(),
		Command:   message.CMD_ERROR.Enum(),
		AuthorId:  proto.Uint32(*msg.AuthorId),
		SessionId: proto.String(*msg.SessionId),
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
		Target:    message.TARGET_CARDS.Enum(),
		Command:   message.CMD_SUCCES.Enum(),
		AuthorId:  proto.Uint32(*msg.AuthorId),
		SessionId: proto.String(*msg.SessionId),
		}
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgCardCreate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgCardUpdate(conn net.Conn, msg *message.Msg) {
	card := &Card{
		0,
		*msg.Cards.Name,
		*msg.Cards.Desc,
		*msg.Cards.ColumnId,
		*msg.Cards.ProjectId,
		msg.Cards.Tags,
		*msg.Cards.UserId,
		msg.Cards.ScriptsIds,
		msg.Cards.Write,
	}
	var answer *message.Msg
	if err := card.Update(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
		Target:    message.TARGET_CARDS.Enum(),
		Command:   message.CMD_ERROR.Enum(),
		AuthorId:  proto.Uint32(*msg.AuthorId),
		SessionId: proto.String(*msg.SessionId),
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
		Target:    message.TARGET_CARDS.Enum(),
		Command:   message.CMD_SUCCES.Enum(),
		AuthorId:  proto.Uint32(*msg.AuthorId),
		SessionId: proto.String(*msg.SessionId),
		}
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgCardUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

// Cette fonction a une gestion synchrone des messages (traitement les uns apres les autres, pas de traitements paralleles)
// Il faut faire une pool de worker, un dispacher et lancer l'operation a effectuer dans le dispatch.
func MsgCard(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgCardCreate(conn, msg)
	case message.CMD_MODIFY:
	case message.CMD_DELETE:
	case message.CMD_GET:
	case message.CMD_MOVE:
	case message.CMD_CONNECT:
	case message.CMD_DISCONNECT:
	case message.CMD_ERROR:
	case message.CMD_NONE:
	}
}
