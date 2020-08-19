package http

type Response struct {
	Error error
	Data  interface{}
}

type BaseResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
