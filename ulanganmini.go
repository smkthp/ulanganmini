package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Welcome to Ulangan Mini!")

	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err == nil {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	conn.Write([]byte("\u001B[2J"))

	conn.Write([]byte("Halo!\n"))

	go func() {
		for {
			conn.Write([]byte(">>> "))
			buffer := bufio.NewReader(conn)
			msg, err := buffer.ReadString('\n')
			if err != nil {
				continue
			}

			if strings.HasPrefix(msg, "!exit") {
				conn.Write([]byte("Bye!\n"))
				conn.Close()
			}
		}
	}()
}
