package httphandler

type JsonSuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
