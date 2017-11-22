package listperson

import (
	"testing"

	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/icrowley/fake"
)

func TestListPersonResponse(t *testing.T) {
	message := fake.Sentence()
	person := model.NewPerson(
		fake.FirstName(),
		fake.Phone(),
	)

	response := &listPersonResponse{
		message,
		[]*model.Person{person},
	}

	if response.GetMessage() != message {
		t.Error("Message mismatch")
	}

	if response.GetPersonList()[0].ID != person.ID {
		t.Error("Person ID mismatch")
	}

	if response.GetPersonList()[0].Name != person.Name {
		t.Error("Person Name mismatch")
	}

	if response.GetPersonList()[0].Phone != person.Phone {
		t.Error("Person Phone mismatch")
	}
}
