package client

import (
	"bufio"
	"fmt"
	"github.com/mosen/zfe/pkg/comms"
	"github.com/mosen/zfe/pkg/encoding"
	"github.com/mosen/zfe/pkg/message"
	"log"
	"net"
	"time"
)

type ZabbixProxy interface {
	ZabbixClient

	GetProxyConfig() error
	Heartbeat() error
	Start(errorCh chan error)
}

type proxy struct {
	addr   *net.TCPAddr
	conn   *net.TCPConn
	logger *log.Logger
	reader *bufio.Reader

	heartbeatFrequency  time.Duration
	configFrequency     time.Duration
	dataSenderFrequency time.Duration

	errorCh chan error
}

func NewProxy(host string) (ZabbixProxy, error) {
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return nil, err
	}
	proxy := &proxy{addr: addr}

	return proxy, nil
}

func (p *proxy) Dial() error {
	conn, err := net.DialTCP("tcp4", nil, p.addr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	p.conn = conn
	return nil
}

func (p *proxy) GetProxyConfig() error {
	fmt.Println("GetProxyConfig()")
	req := message.NewHostRequest(comms.ZbxProtoValueProxyConfig)
	reqData, err := encoding.Encode(req)
	if err != nil {
		return err
	}

	_, err = p.conn.Write(reqData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (p *proxy) Heartbeat() error {
	return nil
}

//func (p *proxy) Read() error {
//	p.reader.Read()
//}

func (p *proxy) Start(errorCh chan error) {
	fmt.Println("Starting proxy")
	if err := p.Dial(); err != nil {
		errorCh <- err
	}

	if err := p.GetProxyConfig(); err != nil {
		errorCh <- err
	}

	p.errorCh = errorCh

	reader := bufio.NewReader(p.conn)
	p.reader = reader

	for {
		select {
		case <-p.errorCh:
			fmt.Println("Proxy is shutting down...")
			return
		default:
			_, _ = encoding.DecodeNext(p.reader)

		}
	}
}

func (p *proxy) Stop() {
	close(p.errorCh)
}
