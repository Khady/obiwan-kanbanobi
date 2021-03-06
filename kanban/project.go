package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

type Project struct {
	Id        uint32
	Name      string
	admins_id []uint32
	Read      []uint32
	Content   string
}

func MsgProjectCreate(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		Name:      *msg.Projects.Name,
		admins_id: msg.Projects.AdminsId,
		Read:      []uint32{*msg.AuthorId},
		Content:   *msg.Projects.Content,
	}
	var answer *message.Msg
	if err := proj.Add(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(31),
			},
		}
	} else {
		proj.Get(dbPool)
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_GET.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Projects: &message.Msg_Projects{
				Id:       &proj.Id,
				Name:     &proj.Name,
				Content:  &proj.Content,
				AdminsId: proj.admins_id,
				Read:     proj.Read},
		}
		notifyUsers(msg)
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectCreate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectUpdate(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		Id:        *msg.Projects.Id,
		Name:      *msg.Projects.Name,
		admins_id: msg.Projects.AdminsId,
		Read:      msg.Projects.Read,
		Content:   *msg.Projects.Content,
	}

	var answer *message.Msg
	if err := proj.Update(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(32),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
		notifyUsers(msg)
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectDelete(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		Id: *msg.Projects.Id,
	}
	var answer *message.Msg
	temp := getListOfUser(proj.Id)
	if err := proj.Del(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(32),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
		sendMsgToallUser(temp, msg)
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectGet(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		Id: *msg.Projects.Id,
	}
	var answer *message.Msg
	if err := proj.Get(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(35),			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_GET.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Projects: &message.Msg_Projects{
				Id:       &proj.Id,
				Name:     &proj.Name,
				Content:  &proj.Content,
				AdminsId: proj.admins_id,
				Read:     proj.Read},
		}
		notifyUsers(msg)
	}
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgProjectUpdate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgProjectGetBoard(conn net.Conn, msg *message.Msg) {
	proj := &Project{
		*msg.Projects.Id,
		*msg.Projects.Name,
		nil,
		nil,
		*msg.Projects.Content,
	}
	var answer *message.Msg

	if err := proj.GetById(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_PROJECTS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(36),
			},
		}
	} else {
		// add verif for read right
		if board, err := proj.GetColumnByProjectId(dbPool); err != nil {
			answer = &message.Msg{
				Target:    message.TARGET_PROJECTS.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(36),
				},
			}
		} else {
			answer = &message.Msg{
				Target:    message.TARGET_PROJECTS.Enum(),
				Command:   message.CMD_SUCCES.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Projects: &message.Msg_Projects{
					Id:             proto.Uint32(proj.Id),
					Name:           &proj.Name,
					Content:        &proj.Content,
					AdminsId:       proj.admins_id,
					Read:           proj.Read,
					ProjectColumns: ConvertTabOfColumnToMessage(board),
				},
			}
		}
	}
	sendKanbanMsg(conn, answer)
}

// modifier pour faire des column
func ConvertTabOfColumnToMessage(p []Column) []*message.Msg_Columns {
	var ret []*message.Msg_Columns

	for n := 0; n < len(p); n++ {
		ret = append(ret, &message.Msg_Columns{
			Id:         proto.Uint32(p[n].Id),
			ProjectId:  proto.Uint32(p[n].Project_id),
			Name:       proto.String(p[n].Name),
			Desc:       proto.String(p[n].Content),
			Tags:       p[n].Tags,
			ScriptsIds: p[n].Scripts_id,
			Write:      p[n].Write,
		})
	}
	return ret
}

func MsgProject(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgProjectCreate(conn, msg)
	case message.CMD_MODIFY:
		MsgProjectUpdate(conn, msg)
	case message.CMD_DELETE:
		MsgProjectDelete(conn, msg)
	case message.CMD_GET:
		MsgProjectGet(conn, msg)
	case message.CMD_MOVE:
		MsgProjectUpdate(conn, msg)
	case message.CMD_GETBOARD:
		MsgProjectGetBoard(conn, msg)
	default:
		UnknowCommand(conn, msg)
	}
}
