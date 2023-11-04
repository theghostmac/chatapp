package main

import (
	"fmt"
	"log"
	"net"
)

const PORT = ":6868"
const SafeMode = true

func main() {
	ln, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("ERROR: could not listen to port, due to: %s\n", err)
	}

	log.Printf("Listening to TCP connections on Port: %s ...\n", PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("ERROR: could not accept a connection: %s, due to: %s\n", conn.RemoteAddr(), err)
		}

		log.Printf("Accepted connection from: %s\n", safeRemoteAddr(conn))

		go handleConnection(conn)
	}
}

func safeRemoteAddr(conn net.Conn) string {
	if SafeMode {
		return "[REDACTED]"
	} else {
		return conn.RemoteAddr().String()
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Connection from %s closed: %s\n", conn.RemoteAddr(), err)
			return
		}

		data := string(buffer[:n])
		fmt.Printf("Received data from %s: %s\n", safeRemoteAddr(conn), data)

		// You can process the data here or send it to an outgoing channel.
		// For example, if you have an outgoing channel, you can do:
		// outgoing <- data
	}
}
