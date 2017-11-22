package repository

import "github.com/GoJogja/go-clean-arch/domain/model"

type PersonRepository interface {
	Create(p *model.Person) bool
	All() []*model.Person
}
