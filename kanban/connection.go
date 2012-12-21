package main

import (
	"bitbucket.org/ongisnotaguild/obi-wan-kanbanobi/protocole"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"net"
	"strconv"
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

func handleConnection(conn net.Conn) {
	header := true
	var size int
	var buf []byte
	defer conn.Close()
	defer CONNECTION_LIST.delConnection(conn)
	defer LOGGER.Print("Connection close")
	for {
		if header {
			buf = make([]byte, 8)
		} else {
			buf = make([]byte, size)
		}
		n, err := conn.Read(buf[0:])
		if err != nil {
			LOGGER.Print("get client data error: ", err)
			return
		}
		if header {
			size, _ = strconv.Atoi(string(buf[0 : n-1]))
			fmt.Println("taille recup", size)
			// size, _ = strconv.Atoi(string(buf[0:n]))
			// fmt.Println("et donc ca fait une size de", size)
			header = false
		} else {
			test := &message.Msg{
				Target:    message.TARGET_IDENT.Enum(),
				Command:   message.CMD_CONNECT.Enum(),
				AuthorId:  proto.Uint32(0),
				SessionId: proto.String(""),
			}
			data, err := proto.Marshal(test)
			if err != nil {
				fmt.Println(err)
			}
			// constituer une eventuelle reponse
			fmt.Fprintf(conn, "%d%s", len(data), data)
			header = true
		}
		fmt.Println("size", size, ", len", n)
		fmt.Printf("%#v\n", buf[0:n])
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
