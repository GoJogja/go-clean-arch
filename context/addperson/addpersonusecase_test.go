package addperson

import (
	"testing"

	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/icrowley/fake"
)

type mockRequest struct {
	name  string
	phone string
}

func (r *mockRequest) GetName() string {
	return r.name
}

func (r *mockRequest) GetPhone() string {
	return r.phone
}

type mockRepo struct {
	persons []*model.Person
}

func (r *mockRepo) Create(p *model.Person) bool {
	r.persons = append(r.persons, p)
	return true
}

func (r *mockRepo) All() []*model.Person {
	return []*model.Person{}
}

func TestAddPerson(t *testing.T) {
	name := fake.FirstName()
	phone := fake.Phone()

	req := &mockRequest{
		name,
		phone,
	}

	repo := &mockRepo{
		make([]*model.Person, 0),
	}

	uc := NewAddPersonUseCase(repo)

	resp := uc.AddPerson(req)

	if resp.GetMessage() != "person Successfully added" {
		t.Error("Create failed")
	}

	if len(repo.persons) == 0 {
		t.Error("Data not inserted")
	}

	if repo.persons[0].Name != name {
		t.Errorf("Expected Name: %s, found %s", name, repo.persons[0].Name)
	}

	if repo.persons[0].Phone != phone {
		t.Errorf("Expected Phone: %s, found %s", phone, repo.persons[0].Phone)
	}
}
