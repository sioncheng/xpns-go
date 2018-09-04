package server

import "net"

//Client represent a remote peer connection
type Client struct {
	TCPConn net.Conn
}

//NewClient create client object
func NewClient(conn net.Conn) *Client {
	c := Client{
		TCPConn: conn,
	}
	return &c
}

//Start start client work
func (p *Client) Start() {

}

//Stop stop client work
func (p *Client) Stop() {

}
