package listperson

import "github.com/GoJogja/go-clean-arch/domain/model"

type ListPersonResponse interface {
	GetMessage() string
	GetPersonList() []*model.Person
}

type listPersonResponse struct {
	message    string
	personList []*model.Person
}

func (r *listPersonResponse) GetMessage() string {
	return r.message
}

func (r *listPersonResponse) GetPersonList() []*model.Person {
	return r.personList
}
