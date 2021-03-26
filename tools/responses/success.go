package responses

import "net/http"

const (
	// StatusText = Success
	StatusText = "Success"
	// AppCode = 200/OK
	AppCode    = http.StatusOK
)

// SuccessResponse to return
type SuccessResponse struct {
	// StatusText string      `json:"status"`
	// AppCode    int64       `json:"code,omitempty"`
	Data       interface{} `json:"tasks"`
}