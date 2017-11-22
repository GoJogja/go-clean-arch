package listperson

import (
	"net/http"

	"github.com/GoJogja/go-clean-arch/context/listperson"
	"github.com/GoJogja/go-clean-arch/transport/httphandler"
)

type ListPersonHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	uc listperson.ListPersonUseCase
	rr httphandler.RequestReader
}

func NewHandler(
	uc listperson.ListPersonUseCase,
	rr httphandler.RequestReader,
) ListPersonHandler {
	return &handler{uc, rr}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	res := h.uc.ListPerson()
	jsonRes := newJsonListPerson(res)
	err := httphandler.JsonOK(
		w,
		jsonRes,
	)

	if err != nil {
		panic(err)
	}
}
