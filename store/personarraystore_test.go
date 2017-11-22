package store

import (
	"testing"

	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/icrowley/fake"
	uuid "github.com/satori/go.uuid"
)

func TestCreatePersonArrayStore(t *testing.T) {
	store := &personArrayStore{
		make(map[uuid.UUID]*model.Person),
	}

	person := model.NewPerson(
		fake.FirstName(),
		fake.Phone(),
	)

	success := store.Create(person)

	if !success {
		t.Error("Failed to create person")
	}

	data := store.personStore[person.GetID()]

	if data.ID != person.ID {
		t.Errorf("Expected ID: %s, found %s", person.ID, data.ID)
	}

	if data.Name != person.Name {
		t.Errorf("Expected Name: %s, found %s", person.Name, data.Name)
	}

	if data.Phone != person.Phone {
		t.Errorf("Expected Phone: %s, found %s", person.Phone, data.Phone)
	}
}

func TestAllPersonArrayStore(t *testing.T) {
	store := &personArrayStore{
		make(map[uuid.UUID]*model.Person),
	}

	person := model.NewPerson(
		fake.FirstName(),
		fake.Phone(),
	)

	success := store.Create(person)

	if !success {
		t.Error("Failed to create person")
	}

	list := store.All()

	if len(list) != 1 {
		t.Error("Expected item count mismatch")
	}

	if list[0].ID != person.ID {
		t.Errorf("Expected ID: %s, found %s", person.ID, list[0].ID)
	}

	if list[0].Name != person.Name {
		t.Errorf("Expected Name: %s, found %s", person.Name, list[0].Name)
	}

	if list[0].Phone != person.Phone {
		t.Errorf("Expected Phone: %s, found %s", person.Phone, list[0].Phone)
	}
}
