package client

import (
	"github.com/mosen/zfe/pkg/message"
	"log"
	"time"
)

type ZabbixProxy interface {
	ZabbixClient

	GetProxyConfig() error
	Heartbeat() error
}

type proxy struct {
	*client
	logger *log.Logger

	heartbeatFrequency  time.Duration
	configFrequency     time.Duration
	dataSenderFrequency time.Duration

	stopCh chan struct{}
}

func NewProxy(host string) (ZabbixProxy, error) {

	//c, err := NewClient(host)
	//if err != nil {
	//	return nil, err
	//}

	return &proxy{}, nil
}

func (p *proxy) GetProxyConfig() error {
	p.logger.Println("GetProxyConfig()")
	req := message.NewHostRequest(message.REQ_PROXY_CONFIG)
	reqData, err := message.Encode(req)
	if err != nil {
		return err
	}

	_, err = p.conn.Write(reqData)
	if err != nil {
		return err
	}

	return nil
}

func (p *proxy) Heartbeat() error {
	return nil
}

func (p *proxy) Start() {
	for {
		select {
		case <-p.stopCh:
			p.logger.Println("Proxy is shutting down...")
			return
		}
	}
}

func (p *proxy) Stop() {
	close(p.stopCh)
}
