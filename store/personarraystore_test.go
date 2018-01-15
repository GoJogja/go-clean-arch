package store

import (
	"testing"

	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/arseto/assertion"
	"github.com/icrowley/fake"
	uuid "github.com/satori/go.uuid"
)

//type personArrayStoreTester struct {
//assert assertion.Assert
//}

//func TestPersonArrayStore(t *testing.T) {
//tester := personArrayStoreTester{assertion.NewAssert(t)}
//tester.testCreate()
//}

//func (tester *personArrayStoreTester) testCreate() {
//store := &personArrayStore{
//make(map[uuid.UUID]*model.Person),
//}

//person := model.NewPerson(
//fake.FirstName(),
//fake.Phone(),
//)

//success := store.Create(person)

//if !success {
//t.Error("Failed to create person")
//}

//data := store.personStore[person.GetID()]

//tester.assert.Equals(data.ID, person.ID)
//tester.assert.Equals(data.Name, person.Name)
//tester.assert.Equals(data.Phone, person.Phone)
//}

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

	assert := assertion.NewAssert(t)

	assert.Equals(data.ID, person.ID)
	assert.Equals(data.Name, person.Name)
	assert.Equals(data.Phone, person.Phone)
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

	assert := assertion.NewAssert(t)

	assert.Equals(list[0].ID, person.ID)
	assert.Equals(list[0].Name, person.Name)
	assert.Equals(list[0].Phone, person.Phone)
}
