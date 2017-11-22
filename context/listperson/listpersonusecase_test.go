package listperson

import (
	"testing"

	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/icrowley/fake"
)

type mockRepo struct {
	persons []*model.Person
}

func (r *mockRepo) Create(p *model.Person) bool {
	return true
}

func (r *mockRepo) All() []*model.Person {
	return r.persons
}

func TestListPerson(t *testing.T) {
	person := model.NewPerson(
		fake.FirstName(),
		fake.Phone(),
	)

	repo := &mockRepo{
		[]*model.Person{person},
	}

	uc := NewListPersonUseCase(repo)
	response := uc.ListPerson()

	if response.GetMessage() != "Person list fetched" {
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
