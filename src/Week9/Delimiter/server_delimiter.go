package Delimiter

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func StartWithDelimiter() {
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
		go handleWithDelimiter(c)
	}
}

func handleWithDelimiter(c net.Conn) {
	defer c.Close()
	for {
		reader := bufio.NewReader(c)
		for {
			data, err := reader.ReadSlice('\n')
			if err != nil && err == io.EOF {
				break
			}
			fmt.Println("received msg from client, the msg is ", string(data), "the length of msg is ", len(data))
		}
	}
}
