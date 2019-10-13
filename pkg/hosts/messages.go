package hosts

type createRequest struct {
}

type indexResponse struct {
	Data   []Host   `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}
