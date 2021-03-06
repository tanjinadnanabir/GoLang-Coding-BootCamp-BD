package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	nl, err := net.Listen("tcp", ":8888") // localhost:port, 127.0.0.1:port, my_ip:port
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr)

	conn.Write([]byte("Welcome to my world!"))

	conn.Close()
	nl.Close()
}
