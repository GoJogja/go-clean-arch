package model

import "strings"

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{
		strings.Title(strings.TrimSpace(name)),
	}
}

func (h *Person) GetName() string {
	return h.name
}
