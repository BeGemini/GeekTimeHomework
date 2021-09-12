package Week9

import (
	"fmt"
	"io"
	"net"
)

const (
	DEFINED_LENGTH = 1000
)

func StartWithFixLength() {
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
		go handleConnWithFixLength(c)
	}
}

func handleConnWithFixLength(c net.Conn) {
	defer c.Close()
	for {
		var data = make([]byte, DEFINED_LENGTH)
		n, err := c.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println("Error occurred during reading from connection,", err)
		}
		if n > 0 {
			fmt.Println("received msg from client, the msg is ", string(data[:n]), "the length of msg is ", n)
		}
	}
}
