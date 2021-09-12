package FixLength

import (
	"log"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestFixLength(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= 3; i++ {
		msg1 := []byte(strconv.Itoa(i) + "aaaa")
		msg2 := []byte(strconv.Itoa(i) + "bbbbb")
		fixLen1 := DEFINED_LENGTH - len(msg1)
		fixLen2 := DEFINED_LENGTH - len(msg2)
		msg1 = append(msg1, make([]byte, fixLen1)...)
		msg2 = append(msg2, make([]byte, fixLen2)...)
		conn.Write(msg1)
		conn.Write(msg2)
	}
	time.Sleep(time.Second)
}
