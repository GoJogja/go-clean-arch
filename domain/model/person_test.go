package model

import "testing"

func TestNewPerson(t *testing.T) {
	input := "Arseto"
	expected := "Arseto"

	person := NewPerson(input)
	if person.GetName() != expected {
		t.Errorf("Expected name %s, found %s", expected, person.GetName())
	}
}

func TestNewPersonLowerCase(t *testing.T) {
	input := "arseto"
	expected := "Arseto"

	person := NewPerson(input)
	if person.GetName() != expected {
		t.Errorf("Expected name %s, found %s", expected, person.GetName())
	}
}

func TestNewPersonWithWhitespace(t *testing.T) {
	input := " arseto "
	expected := "Arseto"

	person := NewPerson(input)
	if person.GetName() != expected {
		t.Errorf("Expected name %s, found %s", expected, person.GetName())
	}
}
