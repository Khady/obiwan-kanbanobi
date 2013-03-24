package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

func MsgAdminCreate(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
		true,
	}
	var answer *message.Msg
	if user.HaveRight((*msg.AuthorId)) {
		answer = &message.Msg{
			Target:    message.TARGET_ADMIN.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2),
			},
		}
	} else if err := user.PutAdmin(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_ADMIN.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(17),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_ADMIN.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgAdminDelete(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
		true,
	}
	var answer *message.Msg

	if user.HaveRight((*msg.AuthorId)) {
		answer = &message.Msg{
			Target:    message.TARGET_ADMIN.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2),
			},
		}
	} else if err := user.Unadmin(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_ADMIN.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(18),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_ADMIN.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgAdminGet(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
		true,
	}

	var answer *message.Msg

	if user.Id != 0 {
		if err := user.GetById(dbPool); err != nil {
			answer = &message.Msg{
				Target:    message.TARGET_ADMIN.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(15),
				},
			}
		} else {
			answer = &message.Msg{
				Target:    message.TARGET_ADMIN.Enum(),
				Command:   message.CMD_SUCCES.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Users: &message.Msg_Users{
					Id:    proto.Uint32(user.Id),
					Name:  &user.Name,
					Admin: &user.Admin,
					Mail:  &user.Mail,
				},
			}
		}
	} else {
		if err := user.GetByName(dbPool); err != nil {
			answer = &message.Msg{
				Target:    message.TARGET_ADMIN.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(15),
				},
			}
		} else {
			answer = &message.Msg{
				Target:    message.TARGET_ADMIN.Enum(),
				Command:   message.CMD_SUCCES.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Users: &message.Msg_Users{
					Id:    proto.Uint32(user.Id),
					Name:  &user.Name,
					Admin: &user.Admin,
					Mail:  &user.Mail,
				},
			}
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgAdmin(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgAdminCreate(conn, msg)
	case message.CMD_DELETE:
		MsgAdminDelete(conn, msg)
	case message.CMD_GET:
		MsgAdminGet(conn, msg)
	}
}
