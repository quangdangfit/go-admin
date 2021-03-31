package schema

type BaseResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
