package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

type Column struct {
	Id         uint32
	Name       string
	Project_id uint32
	Content    string
	Tags       []string
	Scripts_id []uint32
	Write      []uint32
}

// Il y a moyen de factoriser beaucoup le code des fonctions ici.
// Il faut juste penser a faire une gestion d'erreur un peu avant la fonction
// commune pour renvoyer le bon code d'erreur (verifier que l'ID d'une carte
// existe avant de faire un delete par exemple)

func MsgColumnCheckDefaultDesc(desc *string) *string {
	var finalDesc *string
	defaultDesc := "New column"
	if desc != nil {
		finalDesc = desc
	} else {
		finalDesc = &defaultDesc
	}
	return finalDesc
}

func MsgColumnCheckDefaultTags(tags *[]string) *[]string {
	var finalTags *[]string
	defaultTags := []string{}
	if tags != nil {
		finalTags = tags
	} else {
		finalTags = &defaultTags
	}
	return finalTags
}

func MsgColumnCheckDefaultScriptsId(scriptsId *[]uint32) *[]uint32 {
	var finalScriptsId *[]uint32
	defaultScriptsId := []uint32{}
	if scriptsId != nil {
		finalScriptsId = scriptsId
	} else {
		finalScriptsId = &defaultScriptsId
	}
	return finalScriptsId
}

func MsgColumnCheckDefaultWrite(write *[]uint32) *[]uint32 {
	var finalWrite *[]uint32
	defaultWrite := []uint32{}
	if write != nil {
		finalWrite = write
	} else {
		finalWrite = &defaultWrite
	}
	return finalWrite
}

// msg.Columns.UserId est utilise par defaut pour le moment. Mais c'est un champ optionnel.
// Il faudrait faire un test pour savoir si c'est le author_id ou lui qui est utilise.
func MsgColumnCreate(conn net.Conn, msg *message.Msg) {
	LOGGER.Print("New TARGET_COLUMN msg, CMD = CREATE")
	description := MsgColumnCheckDefaultDesc(msg.Columns.Desc)
	tags := MsgColumnCheckDefaultTags(&msg.Columns.Tags)
	scriptsid := MsgColumnCheckDefaultScriptsId(&msg.Columns.ScriptsIds)
	write := MsgColumnCheckDefaultWrite(&msg.Columns.Write)
	column := &Column{
		0,
		*msg.Columns.Name,
		*msg.Columns.ProjectId,
		*description,
		*tags,
		*scriptsid,
		*write,
		// msg.Columns.Tags,
		// msg.Columns.ScriptsIds,
		// msg.Columns.Write,
	}
	var answer *message.Msg
	if err := column.Add(dbPool); err != nil {
		LOGGER.Print("Impossible to create new column", err)
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_COLUMNS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
			},
		}
	} else {
		LOGGER.Print("New column")
		// Envoyer un message de succes ici
		answer = &message.Msg{
			Target:    message.TARGET_COLUMNS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgColumnUpdate(conn net.Conn, msg *message.Msg) {
	column := &Column{
		*msg.Columns.Id,
		*msg.Columns.Name,
		*msg.Columns.ProjectId,
		*msg.Columns.Desc,
		msg.Columns.Tags,
		msg.Columns.ScriptsIds,
		msg.Columns.Write,
	}
	var answer *message.Msg
	if err := column.Update(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_COLUMNS.Enum(),
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
			Target:    message.TARGET_COLUMNS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
}

// Il faut rajouter un traitement ici pour se debarasser des cartes qui sont dans la column qu'on delete
func MsgColumnDelete(conn net.Conn, msg *message.Msg) {
	column := &Column{
		Id: *msg.Columns.Id,
	}
	var answer *message.Msg
	if err := column.Del(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_COLUMNS.Enum(),
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
			Target:    message.TARGET_COLUMNS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgColumnGet(conn net.Conn, msg *message.Msg) {
	column := &Column{
		Id: *msg.Columns.Id,
	}
	var answer *message.Msg
	if err := column.Get(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_COLUMNS.Enum(),
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
			Target:    message.TARGET_COLUMNS.Enum(),
			Command:   message.CMD_GET.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Columns: &message.Msg_Columns{
				Id:         proto.Uint32(column.Id),
				ProjectId:  proto.Uint32(column.Project_id),
				Name:       proto.String(column.Name),
				Desc:       proto.String(column.Content),
				Tags:       column.Tags,
				ScriptsIds: column.Scripts_id,
				Write:      column.Write,
			},
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgColumnGetBoard(conn net.Conn, msg *message.Msg) {
	col := &Column{
		*msg.Columns.Id,
		*msg.Columns.Name,
		*msg.Columns.ProjectId,
		"",
		nil,
		nil,
		nil,
	}
	var answer *message.Msg

	// add verif for read right
	if board, err := col.GetCardByColumnId(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_COLUMNS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(36),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_COLUMNS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Columns: &message.Msg_Columns{
				ProjectId:   proto.Uint32(col.Project_id),
				Id:          proto.Uint32(col.Id),
				Name:        &col.Name,
				ColumnCards: ConvertTabOfCardToMessage(board),
			},
		}
	}
	sendKanbanMsg(conn, answer)
}

func ConvertTabOfCardToMessage(p []Card) []*message.Msg_Cards {
	var ret []*message.Msg_Cards

	for n := 0; n < len(p); n++ {
		ret = append(ret, &message.Msg_Cards{
			Id:         proto.Uint32(p[n].Id),
			ProjectId:  proto.Uint32(p[n].Project_id),
			ColumnId:   proto.Uint32(p[n].Column_id),
			Name:       proto.String(p[n].Name),
			Desc:       proto.String(p[n].Content),
			Tags:       p[n].Tags,
			UserId:     proto.Uint32(p[n].User_id),
			ScriptsIds: p[n].Scripts_id,
			Write:      p[n].Write,
		})
	}
	return ret
}

// Cette fonction a une gestion synchrone des messages (traitement les uns apres les autres, pas de traitements paralleles)
// Il faut faire une pool de worker, un dispacher et lancer l'operation a effectuer dans le dispatch.
func MsgColumn(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgColumnCreate(conn, msg)
	case message.CMD_MODIFY:
		MsgColumnUpdate(conn, msg)
	case message.CMD_DELETE:
		MsgColumnDelete(conn, msg)
	case message.CMD_GET:
		MsgColumnGet(conn, msg)
	case message.CMD_GETBOARD:
		MsgColumnGetBoard(conn, msg)
	case message.CMD_MOVE:
		MsgColumnUpdate(conn, msg)
	default:
		UnknowCommand(conn, msg)
	}
}
