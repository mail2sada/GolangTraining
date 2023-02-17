package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	fmt.Println("Demo: tcp Client....")

	conn, err := net.Dial("tcp", "10.10.34.17:3344")

	if err != nil {

		log.Fatal("Could not connect ot server...")
	}

	data := ""

	for {

		fmt.Fscanln(conn, &data)

		fmt.Println(data)

		conn.Write([]byte("This is from client"))

	}
}
