package models

// swagger:response StatusCode
type StatusCode struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}
