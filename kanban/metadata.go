package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

type Metadata struct {
	Id          int
	Object_type int
	Object_id   int
	Data_key    string
	Data_value  string
}

func MsgMetadataCreate(conn net.Conn, msg *message.Msg) {
	m := &Metadata{
	0,
	0,
	0,
	msg.Msg_Metadata.ObjectId,
	msg.Msg_Metadata.Data,
	msg.Msg_Metadata.DataValue }
        var answer *message.Msg
    if err := m.Add(dbPool); err != nil {
        
	answer = &message.Msg{
            Target:    message.TARGET_METADATA.Enum(),
            Command:   message.CMD_ERROR.Enum(),
            AuthorId:  proto.Uint32(*msg.AuthorId),
            SessionId: proto.String(*msg.SessionId),
        
	Error: &message.Msg_Error{
	    ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
	    },
	}
    } else {
        // Envoyer un message de succes ici
	answer = &message.Msg{
	    Target:    message.TARGET_METADATA.Enum(),
	    Command:   message.CMD_SUCCES.Enum(),
	    AuthorId:  proto.Uint32(*msg.AuthorId),
	    SessionId: proto.String(*msg.SessionId),
	}
    }
    if err != nil {
        LOGGER.Print("Error in MsgMetadataCreate(", err, answer)
        return
    }
    conn.Write(write_int32(int32(len(data))))
    conn.Write(data)
}

func MsgMetadataUpdate(conn net.Conn, msg *message.Msg) {
    m := &Metadata{
	0,
	0,
	0,
	msg.Msg_Metadata.ObjectId,
	msg.Msg_Metadata.DataValue }
    
    var answer *message.Msg
	if err := m.Update(dbPool); err != nil {
		// Envoyer un message d'erreur ici
	answer = &message.Msg{
	    Target:    message.TARGET_METADATA.Enum(),
	    Command:   message.CMD_ERROR.Enum(),
	    AuthorId:  proto.Uint32(*msg.AuthorId),
	    SessionId: proto.String(*msg.SessionId),
	    Error: &message.Msg_Error{
	    ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
	    },
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

func MsgMetadataDelete(conn net.Conn, msg *message.Msg) {
    m := &Metadata{
	Id: *msg.Msg_Metadata.Id,
    }
	var answer *message.Msg
    if err := m.Del(dbPool); err != nil {
	// Envoyer un message d'erreur ici
	answer = &message.Msg{
	    Target:    message.TARGET_METADATA.Enum(),
	    Command:   message.CMD_ERROR.Enum(),
	    AuthorId:  proto.Uint32(*msg.AuthorId),
	    SessionId: proto.String(*msg.SessionId),
	Error: &message.Msg_Error{
	    ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
	    },
	}
    } else {
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

func MsgMetadatadGet(conn net.Conn, msg *message.Msg) {
	card := &Card{
		Id: *msg.Cards.Id,
	}
	var answer *message.Msg
	if err := card.Get(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_METADATA.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_CARDS.Enum(),
			Command:   message.CMD_GET.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Cards: &message.Metadata{
				Id:         proto.Uint32(card.Id),
				ProjectId:  proto.Uint32(card.Project_id),
				ColumnId:   proto.Uint32(card.Column_id),
				Name:       proto.String(card.Name),
				Desc:       proto.String(card.Content),
				Tags:       card.Tags,
				UserId:     proto.Uint32(card.User_id),
				ScriptsIds: card.Scripts_id,
				Write:      card.Write,
			},
		}
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgMetadataUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}


func MsgMetadata(conn net.Conn, msg *message.Msg) {
    switch *msg.Command {
    case message.CMD_CREATE:
	MsgCardCreate(conn, msg)
	case message.CMD_MODIFY:
		MsgCardUpdate(conn, msg)
	case message.CMD_DELETE:
		MsgCardDelete(conn, msg)
	case message.CMD_GET:
		MsgCardGet(conn, msg)
	case message.CMD_MOVE:
		MsgCardUpdate(conn, msg)
	}
}
