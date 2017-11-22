package addperson

import (
	"testing"

	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/icrowley/fake"
)

func TestGetMessage(t *testing.T) {
	expected := fake.Sentence()
	resp := &addPersonResponse{
		expected,
		nil,
	}
	if resp.GetMessage() != expected {
		t.Errorf("Expected message %s, found %s", expected, resp.GetMessage())
	}
}

func TestGetPerson(t *testing.T) {
	expected := fake.Sentence()
	person := model.NewPerson(
		fake.FullName(),
		fake.Phone(),
	)
	resp := &addPersonResponse{
		expected,
		person,
	}
	if resp.GetPerson() != person {
		t.Error("Person mismatch")
	}
}
