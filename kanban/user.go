package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"net"
)

type User struct {
	Id       uint32
	Name     string
	Admin    bool
	Password string
	Mail     string
	Active   bool
}

func (u *User) HaveRight(authorId uint32) bool {
	sId := u.Id
	if authorId != u.Id {
		u.Id = authorId
	}
	if ret, err := u.CheckPassword(dbPool, u.Password); err == nil && ret == true {
		if admin, err := u.GetAdminById(dbPool, authorId); (err == nil && admin == true) ||
			authorId == sId {
			u.Id = sId
			return true
		}
	}
	u.Id = sId
	return false
}

func MsgUserCreate(conn net.Conn, msg *message.Msg) {
	var answer *message.Msg
	if msg.Users.Password == nil || msg.Users.Mail == nil {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(11),
			},
		}
	} else {
		user := &User{
			0,
			*msg.Users.Name,
			*msg.Users.Admin,
			*msg.Users.Password,
			*msg.Users.Mail,
			true,
		}
		verifExisting := &User{
			0,
			*msg.Users.Name,
			*msg.Users.Admin,
			*msg.Users.Password,
			*msg.Users.Mail,
			true,
		}
		if err := verifExisting.GetByName(dbPool); user.HaveRight((*msg.AuthorId)) &&
			err != nil && verifExisting.Id == 0 && user.VerifExistingMail(dbPool) == false {
			user.Password = user.Name
			if err := user.Add(dbPool); err != nil {
				answer = &message.Msg{
					Target:    message.TARGET_USERS.Enum(),
					Command:   message.CMD_ERROR.Enum(),
					AuthorId:  proto.Uint32(*msg.AuthorId),
					SessionId: proto.String(*msg.SessionId),
					Error: &message.Msg_Error{
						ErrorId: proto.Uint32(11),
					},
				}
			} else {
				answer = &message.Msg{
					Target:    message.TARGET_USERS.Enum(),
					Command:   message.CMD_SUCCES.Enum(),
					AuthorId:  proto.Uint32(*msg.AuthorId),
					SessionId: proto.String(*msg.SessionId),
				}
			}
		} else {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(2),
				},
			}
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgUserUpdate(conn net.Conn, msg *message.Msg) {

	var answer *message.Msg
	if msg.Users.Password == nil || msg.Users.Mail == nil {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(12),
			},
		}
	} else {
		user := &User{
			*msg.Users.Id,
			*msg.Users.Name,
			*msg.Users.Admin,
			*msg.Users.Password,
			*msg.Users.Mail,
			true,
		}

		if user.HaveRight((*msg.AuthorId)) == false {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(2),
				},
			}
		} else if err := user.Update(dbPool); err != nil {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(12),
				},
			}
		} else {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
				Command:   message.CMD_SUCCES.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
			}
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgUserPassword(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Password.Id,
		"",
		false,
		*msg.Password.Oldpassword,
		"",
		true,
	}

	var answer *message.Msg

	if user.HaveRight((*msg.AuthorId)) == false {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2),
			},
		}
	} else if err := user.ChangePassword(dbPool, *msg.Password.Newpassword); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(13),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgUserDelete(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		"",
		"",
		true,
	}

	var answer *message.Msg
	if user.HaveRight((*msg.AuthorId)) == false {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(2),
			},
		}
	} else if err := user.Del(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(14),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
		CONNECTION_LIST.del(*msg.AuthorId)
		session := &Session{
			User_uid: *msg.AuthorId,
		}
		session.DelByUserId(dbPool)
	}
	sendKanbanMsg(conn, answer)
}

func MsgUserGet(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		"",
		"",
		true,
	}

	var answer *message.Msg

	if user.Id != 0 {
		if err := user.GetById(dbPool); err != nil {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(15),
				},
			}
		} else {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
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
				Target:    message.TARGET_USERS.Enum(),
				Command:   message.CMD_ERROR.Enum(),
				AuthorId:  proto.Uint32(*msg.AuthorId),
				SessionId: proto.String(*msg.SessionId),
				Error: &message.Msg_Error{
					ErrorId: proto.Uint32(15),
				},
			}
		} else {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
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

func MsgUserGetBoard(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		"",
		"",
		true,
	}
	var answer *message.Msg

	if board, err := user.GetProjectByUserId(dbPool); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(15),
			},
		}
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Users: &message.Msg_Users{
				Id:          proto.Uint32(user.Id),
				Name:        &user.Name,
				Admin:       &user.Admin,
				UserProject: ConvertTabOfProjectToMessage(board),
			},
		}
	}
	sendKanbanMsg(conn, answer)
}

func ConvertTabOfProjectToMessage(p []Project) []*message.Msg_Projects {
	var ret []*message.Msg_Projects

	for n := 0; n < len(p); n++ {
		ret = append(ret, &message.Msg_Projects{
			Id:      proto.Uint32(p[n].Id),
			Name:    proto.String(p[n].Name),
			Content: proto.String(p[n].Content),
			Read: p[n].Read,
			AdminsId: p[n].admins_id,
		})
	}
	return ret
}

func MsgUser(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgUserCreate(conn, msg)
	case message.CMD_MODIFY:
		MsgUserUpdate(conn, msg)
	case message.CMD_PASSWORD:
		MsgUserPassword(conn, msg)
	case message.CMD_DELETE:
		MsgUserDelete(conn, msg)
	case message.CMD_GET:
		MsgUserGet(conn, msg)
	case message.CMD_GETBOARD:
		MsgUserGetBoard(conn, msg)
	default:
		UnknowCommand(conn, msg)
	}
}
