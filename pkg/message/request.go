package message

const REQ_ACTIVE_CHECKS = "active_checks"
const REQ_PROXY_CONFIG = "proxy config"

type Request struct {
	Request string `json:"request"`
}

type HostRequest struct {
	*Request
	Host    string `json:"host"`
	Version string `json:"version"`
}

func NewHostRequest(request string) *HostRequest {
	hr := &HostRequest{Version: "4.0.0", Host: ""}
	hr.Request.Request = request
	return hr
}

type Response struct {
	Response string `json:"response"`
}
