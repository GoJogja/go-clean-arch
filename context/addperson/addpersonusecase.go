package addperson

import (
	"github.com/GoJogja/go-clean-arch/domain/model"
	"github.com/GoJogja/go-clean-arch/domain/repository"
)

type AddPersonUseCase interface {
	AddPerson(req AddPersonRequest) AddPersonResponse
}

type AddPersonRequest interface {
	GetName() string
	GetPhone() string
}

type useCase struct {
	repo repository.PersonRepository
}

func NewAddPersonUseCase(repo repository.PersonRepository) AddPersonUseCase {
	return &useCase{repo}
}

func (h *useCase) AddPerson(req AddPersonRequest) AddPersonResponse {
	person := model.NewPerson(
		req.GetName(),
		req.GetPhone(),
	)
	success := h.repo.Create(person)

	if success {
		return &addPersonResponse{
			"person Successfully added",
			person,
		}
	}
	return &addPersonResponse{
		"Failed to add person",
		person,
	}
}
