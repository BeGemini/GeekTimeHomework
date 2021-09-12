package LengthField

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func StartLengthField() {
	l, err := net.Listen("tcp", "127.0.0.1:8866")
	if err != nil {
		fmt.Println("Error occurred during listening.", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error occurred during accepting.", err)
		}
		go handleWithLengthFiled(c)
	}
}

func handleWithLengthFiled(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		data, err := reader.ReadSlice('\n')
		if err != nil && err == io.EOF {
			break
		}
		msg := strings.ReplaceAll(string(data), "\n", "")
		length, _ := strconv.Atoi(msg)
		var buf = make([]byte, length)
		n, err := reader.Read(buf)
		fmt.Println("received msg from client, the msg is ", string(buf), "the length of msg is ", n)
	}
}
