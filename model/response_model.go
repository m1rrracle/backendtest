package model

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
}
