package LengthField

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestLengthField(t *testing.T) {
	c, err := net.Dial("tcp", "127.0.0.1:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for i := 1; i <= 3; i++ {
		msg1 := []byte("4\naaaa")
		msg2 := []byte("5\nbbbbb")
		c.Write(msg1)
		c.Write(msg2)
	}
	time.Sleep(time.Second)
}
