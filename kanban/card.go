package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
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
	card = card
}

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
