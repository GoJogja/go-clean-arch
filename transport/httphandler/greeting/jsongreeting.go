package greeting

import (
	"github.com/arseto/go-clean-arch/context/greeting"
	"github.com/arseto/go-clean-arch/transport/httphandler"
)

type jsonGreeting struct {
	Greeting string `json:"greeting"`
}

func newJsonGreeting(
	respModel greeting.GreetingResponse,
	message string,
) *httphandler.JsonSuccessResponse {
	return &httphandler.JsonSuccessResponse{
		&jsonGreeting{
			respModel.GetGreeting(),
		},
		message,
	}
}
