package Week9

import (
	"log"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestTcp(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= 3; i++ {
		msg1 := []byte(strconv.Itoa(i) + "aaaa")
		msg2 := []byte(strconv.Itoa(i) + "bbbbb")
		conn.Write(msg1)
		conn.Write(msg2)
	}
	time.Sleep(time.Second)
}
