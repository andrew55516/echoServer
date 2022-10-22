package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	wrter := bufio.NewWriter(conn)

	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error", err)
			break
		}
		log.Printf("Recieved %d bytes: %s\n ", len(s), s)

		log.Println("Write data")
		if _, err := wrter.WriteString(s); err != nil {
			log.Fatalln("Unable to write data", err)
		}
	}
}

func main() {
	listner, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")

	for {
		conn, err := listner.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection", err)
		}
		go echo(conn)
	}
}
