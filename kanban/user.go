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
	if ret, err := u.CheckPassword(dbPool, u.Password); err == nil && ret == true {
		if admin, err := u.GetAdminById(dbPool, authorId); (err == nil && admin == true) || 
			authorId == u.Id {
			return true
		}
	}
	return false
}

func MsgUserCreate(conn net.Conn, msg *message.Msg) {

	user := &User{
		0,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
		true,
	}
	var answer *message.Msg
	if user.HaveRight((*msg.AuthorId)) {
		if err := user.Add(dbPool); err != nil {
			answer = &message.Msg{
				Target:    message.TARGET_USERS.Enum(),
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
	} else {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
			Error: &message.Msg_Error{
				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
			},
		}
	}
	sendKanbanMsg(conn, answer)
}

func MsgUserUpdate(conn net.Conn, msg *message.Msg) {
	// verif if user is admin or if user modifying himself
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
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
				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
			},
		}
	} else if err := user.Update(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
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
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
}

// verifier que ca marche.
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
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_ERROR.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		Error: &message.Msg_Error{
			ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
			},
		}
	} else if err := user.ChangePassword(dbPool, *msg.Password.Newpassword); err != nil {
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
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
		*msg.Users.Password,
		*msg.Users.Mail,
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
				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
			},
		}
	} else if err := user.Del(dbPool); err != nil {
		// Envoyer un message d'erreur ici
		answer = &message.Msg{
			Target:    message.TARGET_USERS.Enum(),
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
			Target:    message.TARGET_USERS.Enum(),
			Command:   message.CMD_SUCCES.Enum(),
			AuthorId:  proto.Uint32(*msg.AuthorId),
			SessionId: proto.String(*msg.SessionId),
		}
	}
	sendKanbanMsg(conn, answer)
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
	}
}
