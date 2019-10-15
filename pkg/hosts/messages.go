package hosts

type createRequest struct {
}

type hostsIndexResponse struct {
	Data   []Host   `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

type templatesIndexResponse struct {
	Data   []Template `json:"data,omitempty"`
	Errors []string   `json:"errors,omitempty"`
}
