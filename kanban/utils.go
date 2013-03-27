package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"net"
	"strconv"
)

func SString_of_SUInt32(s []uint32) []string {
	sstring := make([]string, len(s))
	for index, elem := range s {
		sstring[index] = strconv.FormatUint(uint64(elem), 10)
	}
	return sstring
}

func SUInt32_of_SString(s []string) []uint32 {
	suint := make([]uint32, len(s))
	var res uint64
	for index, elem := range s {
		res, _ = strconv.ParseUint(elem, 10, 0)
		suint[index] = uint32(res)
	}
	return suint
}

// Lecture d'un int32 depuis du binaire. Permet de recuperer la taille de la structure a lire sur le reseau
func read_int32(data []byte) (ret int32, err error) {
	buf := bytes.NewBuffer(data)
	err = binary.Read(buf, binary.BigEndian, &ret)
	return
}

// Ecriture binaire d'un int32 pour l'envoyer sur le reseau.
func write_int32(nb int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, nb)
	return buf.Bytes()
}

func UnknowCommand(conn net.Conn, msg *message.Msg) {
	var answer *message.Msg

	answer = &message.Msg{
		Target:    msg.Target,
		Command:   message.CMD_ERROR.Enum(),
		AuthorId:  proto.Uint32(*msg.AuthorId),
		SessionId: proto.String(*msg.SessionId),
		Error: &message.Msg_Error{
			ErrorId: proto.Uint32(3),
		},
	}
	sendKanbanMsg(conn, answer)
}
