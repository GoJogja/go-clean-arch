package model

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestNewPerson(t *testing.T) {
	input := "Arseto"
	expected := "Arseto"

	phone := fake.Phone()
	person := NewPerson(input, phone)

	if person.GetName() != expected {
		t.Errorf("Expected name %s, found %s", expected, person.GetName())
	}

	if person.GetPhone() != phone {
		t.Errorf("Expected phone %s, found %s", phone, person.GetPhone())
	}
}

func TestNewPersonLowerCase(t *testing.T) {
	input := "arseto"
	expected := "Arseto"

	phone := fake.Phone()
	person := NewPerson(input, phone)

	if person.GetName() != expected {
		t.Errorf("Expected name %s, found %s", expected, person.GetName())
	}
	if person.GetPhone() != phone {
		t.Errorf("Expected phone %s, found %s", phone, person.GetPhone())
	}
}

func TestNewPersonWithWhitespace(t *testing.T) {
	input := " arseto "
	expected := "Arseto"

	phone := fake.Phone()
	person := NewPerson(input, phone)

	if person.GetName() != expected {
		t.Errorf("Expected name %s, found %s", expected, person.GetName())
	}
	if person.GetPhone() != phone {
		t.Errorf("Expected phone %s, found %s", phone, person.GetPhone())
	}
}
