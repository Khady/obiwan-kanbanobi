package main

import (
	"net"
	"fmt"
	"strconv"
)

func handleConnection(conn net.Conn) {
	header := true
	var size int
	var buf []byte
	for {
		if header {
			buf = make([]byte, 4)
		} else {
			buf = make([]byte, size)
		}
		n, err := conn.Read(buf[0:])
		if err != nil {
			LOGGER.Print("get client data error: ", err)
		}
		if header {
			size, _ = strconv.Atoi(string(buf))
			fmt.Println("taille recup", size)
			// size, _ = strconv.Atoi(string(buf[0:n]))
			// fmt.Println("et donc ca fait une size de", size)
			header = false
		} else {
			// constituer une eventuelle reponse
			fmt.Fprintf(conn, strconv.Itoa(n))
			fmt.Fprintf(conn, string(buf[0:n]))
			header = true
		}
		fmt.Println(size)
		fmt.Printf("%#v\n", buf)
	}
	conn.Close()
	LOGGER.Print("Connection close")
}

func startServer() error {
	LOGGER.Print("Launching the server...")
	defer LOGGER.Print("Server quit")
	server_port := ":" + *SPORT
	LOGGER.Printf("Listening on port %s", server_port)
	ln, err := net.Listen("tcp", server_port)
	if err != nil {
		LOGGER.Printf("Impossible to open the server on port %s", server_port)
		return err
	}
	for {
		// garder une liste des connexions ici
		conn, err := ln.Accept()
		if err != nil {
			LOGGER.Print("get client connection error: ", err)
		}
		LOGGER.Print("New client connection, creating new goroutine")
		go handleConnection(conn)
	}
	return err
}