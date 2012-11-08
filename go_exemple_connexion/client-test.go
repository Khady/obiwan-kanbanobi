package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":6010")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Fprintf(conn, "hello server\n")

		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", data)
	}

}
