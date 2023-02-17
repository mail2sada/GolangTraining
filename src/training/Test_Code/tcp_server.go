package main

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {

	fmt.Fprintln(conn, "Welcome to tcp server")
	var data string
	fmt.Fscanln(conn, &data)

	fmt.Printf(data)

}

func main() {

	fmt.Println("Welcome to our tcp server....")

	ln, err := net.Listen("tcp", "10.10.34.17:3344")

	if err != nil {
		log.Fatalln("Listen failed can not continue....")
	}

	for {
		con, acceptErr := ln.Accept()

		if acceptErr != nil {
			log.Println("Could not accept connection")
		}

		go handleConnection(con)

	}
}
