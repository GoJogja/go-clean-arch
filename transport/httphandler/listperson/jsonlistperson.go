package listperson

import (
	"github.com/GoJogja/go-clean-arch/context/listperson"
	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/GoJogja/go-clean-arch/transport/httphandler"
)

type jsonListPerson struct {
	Persons []*model.Person `json:"persons"`
}

func newJsonListPerson(
	respModel listperson.ListPersonResponse,
) *httphandler.JsonSuccessResponse {
	return &httphandler.JsonSuccessResponse{
		&jsonListPerson{
			respModel.GetPersonList(),
		},
		respModel.GetMessage(),
	}
}
