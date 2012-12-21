package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/protocole"
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func (c *connectionList) addConnection(conn net.Conn) {
	CONNECTION_LIST = append(CONNECTION_LIST, conn)
}

func (c *connectionList) delConnection(conn net.Conn) {
	for i, value := range CONNECTION_LIST {
		if value == conn {
			copy(CONNECTION_LIST[i:], CONNECTION_LIST[i+1:])
			CONNECTION_LIST[len(CONNECTION_LIST)-1] = nil
			CONNECTION_LIST = CONNECTION_LIST[:len(CONNECTION_LIST)-1]
		}
	}
}

func read_int32(data []byte) (ret int32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &ret)
	return
}

func testReponse(conn net.Conn) {
	test := &message.Msg{
		Target:    message.TARGET_IDENT.Enum(),
		Command:   message.CMD_CONNECT.Enum(),
		AuthorId:  proto.Uint32(1),
		SessionId: proto.String("superchainedesession"),
	}
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(conn, "%d%s", len(data), data)
}

func readMsg(conn net.Conn, msg []byte, length int) {
	data := &message.Msg{}
	err := proto.Unmarshal(msg[0:length], data)
	if err != nil {
		LOGGER.Print("Impossible to unmarshal the message", msg[0:length])
		return
	}
	switch *data.Target {
	case message.TARGET_USERS:
		LOGGER.Print("read TARGET_USERS message")
	case message.TARGET_COLUMNS:
		LOGGER.Print("read TARGET_COLUMNS message")
	case message.TARGET_PROJECTS:
		LOGGER.Print("read TARGET_PROJECTS message")
	case message.TARGET_CARDS:
		LOGGER.Print("read TARGET_CARDS message")
	case message.TARGET_ADMIN:
		LOGGER.Print("read TARGET_ADMIN message")
	case message.TARGET_IDENT:
		LOGGER.Print("read TARGET_IDENT message")
		testReponse(conn)
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
	defer CONNECTION_LIST.delConnection(conn)
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
			size = int(read_int32(buf))
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
		CONNECTION_LIST.addConnection(conn)
		fmt.Println(CONNECTION_LIST)
		go handleConnection(conn)
	}
	return err
}
