package client

import "net"

type ZabbixClient interface {
	Dial() error
}

type client struct {
	addr *net.TCPAddr
	conn *net.TCPConn
}

func (c *client) Dial() error {
	conn, err := net.DialTCP("tcp4", nil, c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func NewClient(host string) (ZabbixClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return nil, err
	}
	return &client{addr: addr}, nil
}
