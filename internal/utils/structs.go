package utils

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
