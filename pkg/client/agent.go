package client

type ZabbixAgent interface {
	ZabbixClient
}

type agent struct {
	*client
}

func NewAgent() (ZabbixAgent, error) {
	return &agent{}, nil
}
