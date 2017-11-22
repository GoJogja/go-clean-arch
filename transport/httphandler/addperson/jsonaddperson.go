package addperson

import (
	"github.com/GoJogja/go-clean-arch/context/addperson"
	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/GoJogja/go-clean-arch/transport/httphandler"
)

type jsonAddPerson struct {
	Person *model.Person `json:"person"`
}

func newJsonAddPerson(
	respModel addperson.AddPersonResponse,
) *httphandler.JsonSuccessResponse {
	return &httphandler.JsonSuccessResponse{
		&jsonAddPerson{
			respModel.GetPerson(),
		},
		respModel.GetMessage(),
	}
}
