package message

type Item struct {
	Key         string `json:"key"`
	Delay       int    `json:"delay"`
	LastLogSize int    `json:"lastlogsize"`
	Mtime       int    `json:"mtime"`
}

type Metric struct {
	Host  string `json:"host"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Clock int    `json:"clock"`
	Ns    int    `json:"ns"`
}

func NewMetric(host, key, value string) *Metric {
	m := &Metric{Host: host, Key: key, Value: value}
	return m
}

type ActiveCheckResponse struct {
	*Response
	Data []Item `json:"data"`
}

type AgentDataRequest struct {
	*Request
	Clock int `json:"clock"`
	Ns    int `json:"ns"`
}

type AgentDataResponse struct {
	*Response
	Info string `json:"info"`
}
