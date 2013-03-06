package main

type Metadata struct {
	Id          int
	Object_type int
	Object_id   int
	Data_key    string
	Data_value  string
}

func MsgCard(conn net.Conn, msg *message.Msg) {
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
