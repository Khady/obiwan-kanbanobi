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

func MsgUserCreate(conn net.Conn, msg *message.Msg) {
	//verifier si l'user est admin
	user := &User{
		0,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
		true,
	}
	var answer *message.Msg
	if err := user.Add(dbPool); err != nil {
		answer = &message.Msg{
			Target: message.TARGET_USERS.Enum(),
			Command: message.CMD_ERROR.Enum(),
			AuthorId: proto.Uint32(*msg.AuthorId),
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
	data, err := proto.Marshal(answer)
	if err != nil {
		LOGGER.Print("Impossible to marshal msg in MsgUserCreate", err, answer)
		return
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
}

func MsgUserUpdate(conn net.Conn, msg *message.Msg) {
	user := &User{
		*msg.Users.Id,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
		true,
	}
        var answer *message.Msg
	if err := user.Update(dbPool); err != nil {
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
	data, err := proto.Marshal(answer)
        if err != nil {
                LOGGER.Print("Impossible to marshal msg in MsgUserUpdate", err, answer)
                return
        }
        conn.Write(write_int32(int32(len(data))))
        conn.Write(data)
}

// func MsgUserPassword(conn net.Conn, msg *message.Msg) {
// 	user := &User{
// 		*msg.Password.Id,
// 		"",
// 		false,
// 		*msg.Password.Oldpassword,
// 		"",
// 		true,
// 	}
//         var answer *message.Msg
	
// 	if ret, err := user.CheckPassword(dbPool, user.password); err != nil && ret == true {
//                 // Envoyer un message d'erreur ici
// 		if err !=  nil {
// 			answer = &message.Msg{
// 				Target:    message.TARGET_USERS.Enum(),
// 				Command:   message.CMD_ERROR.Enum(),
// 				AuthorId:  proto.Uint32(*msg.AuthorId),
// 				SessionId: proto.String(*msg.SessionId),
// 			Error: &message.Msg_Error{
// 				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
// 				},
// 			}
// 		} else {
// 			answer = &message.Msg{
// 				Target:    message.TARGET_USERS.Enum(),
// 				Command:   message.CMD_ERROR.Enum(),
// 				AuthorId:  proto.Uint32(*msg.AuthorId),
// 				SessionId: proto.String(*msg.SessionId),
// 				Error: &message.Msg_Error{
// 				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
// 				},
// 			}
// 		}
// 	} else {
// 		if err := user.ChangePassword(dbPool, *msg.Password.Newpassword); err != nil {
// 			answer = &message.Msg{
// 				Target:    message.TARGET_USERS.Enum(),
// 				Command:   message.CMD_ERROR.Enum(),
// 				AuthorId:  proto.Uint32(*msg.AuthorId),
// 				SessionId: proto.String(*msg.SessionId),
// 			Error: &message.Msg_Error{
// 				ErrorId: proto.Uint32(1), // remplacer par le vrai code d'erreur ici
// 				},
// 			}
// 		} else {
// 			// Envoyer un message de succes ici                                             
// 			answer = &message.Msg{
// 				Target:    message.TARGET_USERS.Enum(),
// 				Command:   message.CMD_SUCCES.Enum(),
// 				AuthorId:  proto.Uint32(*msg.AuthorId),
// 				SessionId: proto.String(*msg.SessionId),
// 			}
// 		}
// 	}
// 	data, err := proto.Marshal(answer)
//         if err != nil {
//                 LOGGER.Print("Impossible to marshal msg in MsgUserPassword", err, answer)
//                 return
//         }
//         conn.Write(write_int32(int32(len(data))))
//         conn.Write(data)
// }

func MsgUserDelete(conn net.Conn, msg *message.Msg) {
	user := &User{
		0,
		*msg.Users.Name,
		*msg.Users.Admin,
		*msg.Users.Password,
		*msg.Users.Mail,
		true,
	}
	
	var answer *message.Msg
        if err := user.Del(dbPool); err != nil {
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
        data, err := proto.Marshal(answer)
        if err != nil {
                LOGGER.Print("Impossible to marshal msg in MsgUserDelete", err, answer)
                return
        }
        conn.Write(write_int32(int32(len(data))))
        conn.Write(data)
}

func MsgUser(conn net.Conn, msg *message.Msg) {
	switch *msg.Command {
	case message.CMD_CREATE:
		MsgUserCreate(conn, msg)
	case message.CMD_MODIFY:
		MsgUserUpdate(conn, msg)
	case message.CMD_PASSWORD:
//		MsgUserPassword(conn, msg)
	case message.CMD_DELETE:
		MsgUserDelete(conn, msg)
	}
}