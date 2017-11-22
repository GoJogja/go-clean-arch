package model

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Person struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Phone string    `json:"phone"`
}

func NewPerson(name, phone string) *Person {
	return &Person{
		uuid.NewV4(),
		strings.Title(strings.TrimSpace(name)),
		phone,
	}
}

func (p *Person) GetID() uuid.UUID {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetPhone() string {
	return p.Phone
}
