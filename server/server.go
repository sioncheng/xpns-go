package server

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

//Server tcp server for push
type Server struct {
	ListenPort int
	connIn     chan net.Conn
}

//NewServer new server object
func NewServer(listenPort int) *Server {
	s := Server{
		ListenPort: listenPort,
		connIn:     make(chan net.Conn, 10),
	}
	return &s
}

//Start start the tcp push server
func (s *Server) Start() error {
	go s.connReceiver()

	err := s.startListen()
	if err != nil {
		return err
	}

	return nil
}

//Stop stop the tcp push server
func (s *Server) Stop() {

}

func (s *Server) startListen() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.ListenPort))
	if err != nil {
		return err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	for {
		tcpConn, err := tcpListener.Accept()
		if err != nil {
			continue
		}

		log.WithFields(logrus.Fields{"conn": tcpConn.RemoteAddr().String()}).
			Info("client")

		s.connIn <- tcpConn
	}
}

func (s *Server) connReceiver() {
	for {
		select {
		case tcpConn := <-s.connIn:
			client := NewClient(tcpConn)

			log.WithFields(logrus.Fields{"remote_addr": client.TCPConn.RemoteAddr().String()}).
				Info("create new client")
		}
	}
}
