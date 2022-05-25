package servebase

type CommonResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
