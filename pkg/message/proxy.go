package message

type ProxyHeartbeatRequest struct {
	*Request
	Host    string `json:"host"`
	Version string `json:"version"`
}

type ProxyHeartbeatResponse struct {
	*Response
}

type ProxyConfigRequest struct {
	*HostRequest
}

type Table struct {
	Fields []string        `json:"fields"`
	Data   [][]interface{} `json:"data"`
}

type ProxyConfigResponse struct {
	*Request          // not a typo
	GlobalMacro Table `json:"globalmacro"`
	Hosts       Table `json:"hosts"`
	Interface   Table `json:"interface"`
}
