package addperson

import (
	"net/http"

	"github.com/GoJogja/go-clean-arch/context/addperson"
	"github.com/GoJogja/go-clean-arch/transport/httphandler"
)

type AddPersonHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	uc addperson.AddPersonUseCase
	rr httphandler.RequestReader
}

func NewHandler(
	uc addperson.AddPersonUseCase,
	rr httphandler.RequestReader,
) AddPersonHandler {
	return &handler{uc, rr}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {

	req := &addPersonRequest{}
	err := h.rr.GetJsonData(r, req)

	if err != nil {
		//TODO: proper error handling
		panic(err)
	}

	res := h.uc.AddPerson(req)
	jsonRes := newJsonAddPerson(res)

	err = httphandler.JsonOK(
		w,
		jsonRes,
	)

	if err != nil {
		panic(err)
	}
}
