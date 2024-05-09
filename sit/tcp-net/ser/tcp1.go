package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Handle incoming data
	// For example, read data from the connection
	// and send a response back
}

func main() {
	// Start TCP server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	// Set keep-alive parameters
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		// Set keep-alive parameters for the connection
		tcpConn, ok := conn.(*net.TCPConn)
		if ok {
			err := tcpConn.SetKeepAlive(true)
			if err != nil {
				fmt.Println("Error setting keep-alive:", err.Error())
				return
			}
			err = tcpConn.SetKeepAlivePeriod(30 * time.Second)
			if err != nil {
				fmt.Println("Error setting keep-alive period:", err.Error())
				return
			}
			tcpConn.SetNoDelay()
			tcpConn.MultipathTCP()
		}

		// Handle incoming connection in a separate goroutine
		go handleConnection(conn)
	}
}
