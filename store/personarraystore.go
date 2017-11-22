package store

import (
	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/GoJogja/go-clean-arch/domain/repository"
	uuid "github.com/satori/go.uuid"
)

type personArrayStore struct {
	personStore map[uuid.UUID]*model.Person
}

func NewPersonArrayStore() repository.PersonRepository {
	return &personArrayStore{
		make(map[uuid.UUID]*model.Person),
	}
}

func (store *personArrayStore) Create(p *model.Person) bool {
	id := p.GetID()
	store.personStore[id] = p
	return true
}

func (store *personArrayStore) All() (result []*model.Person) {

	result = make([]*model.Person, 0)

	for key, _ := range store.personStore {
		result = append(result, store.personStore[key])
	}
	return
}
