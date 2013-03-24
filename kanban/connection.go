package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/kanban/protocol"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"io"
	"net"
)

func (c *connectionList) add(uid uint32, ulogin string, conn net.Conn) {
	c.ids[uid] = Connection{
		conn,
		uid,
		ulogin,
	}
	_, ok := c.conns[conn]
	if ok {
		c.conns[conn] = append(c.conns[conn], uid)
	} else {
		c.conns[conn] = []uint32{uid}
	}
}

func (c *connectionList) del(uid uint32) {
	conn, ok := c.ids[uid]
	if ok {
		delete(c.ids, uid)
		tab, ok := c.conns[conn.c]
		if ok {
			for i, value := range tab {
				if value == uid {
					copy(tab[i:], tab[i+1:])
					tab = tab[:len(tab)-1]
				}
			}
		}
	}
}

func (c *connectionList) delConn(conn net.Conn) {
	tab, ok := c.conns[conn]
	if ok {
		for _, value := range tab {
			c.del(value)
		}
	}
}

func (c *connectionQueue) add(conn net.Conn) {
	CONNECTION_QUEUE = append(CONNECTION_QUEUE, conn)
}

func (c *connectionQueue) del(conn net.Conn) {
	for i, value := range CONNECTION_QUEUE {
		if value == conn {
			copy(CONNECTION_QUEUE[i:], CONNECTION_QUEUE[i+1:])
			CONNECTION_QUEUE[len(CONNECTION_QUEUE)-1] = nil
			CONNECTION_QUEUE = CONNECTION_QUEUE[:len(CONNECTION_QUEUE)-1]
		}
	}
}

func sendKanbanMsg(conn net.Conn, msg *message.Msg) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	conn.Write(write_int32(int32(len(data))))
	conn.Write(data)
	return nil
}

func unidentifiedUser(conn net.Conn, msg *message.Msg) {
	LOGGER.Print("unidentifiedUser")
	answer := &message.Msg{
		Target:    message.TARGET_IDENT.Enum(),
		Command:   message.CMD_ERROR.Enum(),
		AuthorId:  proto.Uint32(*msg.AuthorId),
		SessionId: proto.String(*msg.SessionId),
		Error: &message.Msg_Error{
			ErrorId: proto.Uint32(1),
		},
	}
	sendKanbanMsg(conn, answer)
}

func readMsg(conn net.Conn, msg []byte, length int) {
	data := &message.Msg{}
	err := proto.Unmarshal(msg[0:length], data)
	if err != nil {
		LOGGER.Print("Impossible to unmarshal the message", msg[0:length])
		return
	}
	if *data.Target != message.TARGET_IDENT && MsgIdentIsUnidentified(conn, data) == false {
		unidentifiedUser(conn, data)
		return
	}
	switch *data.Target {
	case message.TARGET_USERS:
		LOGGER.Print("read TARGET_USERS message")
		MsgUser(conn, data)
	case message.TARGET_COLUMNS:
		LOGGER.Print("read TARGET_COLUMNS message")
		MsgColumn(conn, data)
	case message.TARGET_PROJECTS:
		LOGGER.Print("read TARGET_PROJECTS message")
	case message.TARGET_CARDS:
		LOGGER.Print("read TARGET_CARDS message")
		MsgCard(conn, data)
	case message.TARGET_ADMIN:
		LOGGER.Print("read TARGET_ADMIN message")
		MsgAdmin(conn, data)
	case message.TARGET_IDENT:
		LOGGER.Print("read TARGET_IDENT message")
		MsgIdent(conn, data)
	case message.TARGET_NOTIF:
		LOGGER.Print("read TARGET_NOTIF message")
	case message.TARGET_METADATA:
		LOGGER.Print("read TARGET_METADATA message")
	default:
		LOGGER.Print("Invalid TARGET")
	}
}

func handleConnection(conn net.Conn) {
	header := true
	var size int
	var buf []byte
	defer conn.Close()
	defer CONNECTION_LIST.delConn(conn)
	defer LOGGER.Print("Connection close")
	for {
		if header {
			buf = make([]byte, 4)
		} else {
			buf = make([]byte, size)
		}
		n, err := conn.Read(buf[0:])
		if err == io.EOF {
			return
		}
		if err != nil {
			LOGGER.Print("get client data error: ", err)
		}
		if header {
			tmp_size, err := read_int32(buf)
			if err != nil {
				LOGGER.Print("Impossible to read size", err)
				continue
			}
			size = int(tmp_size) // I should put a check on the size here to raise an error if the size is huge
			header = false
		} else {
			readMsg(conn, buf, n)
			header = true
		}
	}
}

func startServer() error {
	LOGGER.Print("Launching the server...")
	defer LOGGER.Print("Server quit")

	server_port := ":" + *SPORT
	tcpAddr, err := net.ResolveTCPAddr("ip4", server_port)
	if err != nil {
		LOGGER.Printf("The port %s is invalid", server_port)
		return err
	}
	LOGGER.Printf("Listening on port %s", server_port)
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		LOGGER.Printf("Impossible to open the server on port %s", server_port)
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			LOGGER.Print("get client connection error: ", err)
		}
		LOGGER.Printf("New client connection with ip %s, creating new goroutine", conn.RemoteAddr().String())
		CONNECTION_QUEUE.add(conn)
		fmt.Println(CONNECTION_QUEUE)
		go handleConnection(conn)
	}
	return err
}
