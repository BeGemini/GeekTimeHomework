package StickyPacket

import (
	"fmt"
	"io"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var data = make([]byte, 5)
		n, err := c.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println("Error occurred during reading from connection,", err)
		}
		if n > 0 {
			fmt.Println("received msg from client, the msg is ", string(data[:n]), "the length of msg is ", n)
		}
		// write to the connection

	}
}

func StartWithSticky() {
	l, err := net.Listen("tcp", "127.0.0.1:8866")
	if err != nil {
		fmt.Println("Error occurred during listen:", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error occurred during accepting,", err)
		}
		go handleConn(c)
	}
}
