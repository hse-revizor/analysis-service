package models

type ApiError struct {
	Error       string
	Description string `json:"error_description"`
	Code        int    `json:"error_code"`
}
