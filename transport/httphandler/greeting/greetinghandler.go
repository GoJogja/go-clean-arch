package greeting

import (
	"net/http"

	"github.com/arseto/go-clean-arch/context/greeting"
	"github.com/arseto/go-clean-arch/transport/httphandler"
)

type GreetingHandler interface {
	Greet(w http.ResponseWriter, r *http.Request)
}

type greetingHandler struct {
	guc greeting.GreetingUseCase
	rr  httphandler.RequestReader
}

func NewGreetingHandler(
	guc greeting.GreetingUseCase,
	rr httphandler.RequestReader,
) GreetingHandler {
	return &greetingHandler{guc, rr}
}

func (gc *greetingHandler) Greet(w http.ResponseWriter, r *http.Request) {

	req := &greetingRequest{gc.rr.GetRouteParam(r, "name")}
	res := gc.guc.GetGreeting(req)

	jsonRes := newJsonGreeting(res, "Message Loaded")

	err := httphandler.JsonOK(
		w,
		jsonRes,
	)

	if err != nil {
		panic(err)
	}
}
