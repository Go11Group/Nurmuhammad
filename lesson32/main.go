package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var slc []net.Conn

func main() {
	// Listen for incoming connections.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept a connection from a client.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		slc = append(slc, conn)
		go handleConnection(conn) // Handle each connection in a new goroutine.
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		// Read data from the connection.
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Echo the message back to the client.
		fmt.Print("Received message: ", string(message))
		message = message[:len(message)-1]
		Check(conn, message)
	}
}

func Check(conn net.Conn, message string) {
	for _, v := range slc {
		if v != conn {
			fmt.Print("Received message: ", string(message))
			_, err := v.Write([]byte(strings.ToUpper(message) + " FROM SERVER!\n")) // Echo back in uppercase.
			if err != nil {
				fmt.Println("Error writing message:", err)
				return
			}
		}
	}
}
