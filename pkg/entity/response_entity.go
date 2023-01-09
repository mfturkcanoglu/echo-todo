package entity

type RestResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
}
