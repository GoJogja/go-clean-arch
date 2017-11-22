package addperson

import "github.com/GoJogja/go-clean-arch/domain/model"

type AddPersonResponse interface {
	GetMessage() string
	GetPerson() *model.Person
}

type addPersonResponse struct {
	message string
	person  *model.Person
}

func (r *addPersonResponse) GetMessage() string {
	return r.message
}

func (r *addPersonResponse) GetPerson() *model.Person {
	return r.person
}
