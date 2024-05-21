package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Connect to TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return
	}
	defer conn.Close()

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
	}

	// Perform other operations with the TCP connection
}
