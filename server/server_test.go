package server

import (
	"net"
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	s := NewServer(8080)

	go s.Start()

	time.Sleep(200 * time.Millisecond)

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		t.Error(err)
	}

	t.Log(conn.RemoteAddr().String())

	time.Sleep(200 * time.Millisecond)
}
