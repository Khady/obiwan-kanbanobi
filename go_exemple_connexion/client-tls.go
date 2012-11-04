package main

import (
	"bufio"
	"fmt"
	"crypto/tls"
)

func main() {
	conn, err := tls.Dial("tcp", ":6010", nil)
	if err != nil {
		fmt.Println("ca va planter")
		panic(err)
	}
	fmt.Fprintf(conn, "hello server\n")
	client := tls.Client(conn, &tls.Config{})

	err = client.Handshake()
        if err != nil {
		panic(err)
        }

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", data)
}
