package listperson

import (
	"github.com/GoJogja/go-clean-arch/domain/repository"
)

type ListPersonUseCase interface {
	ListPerson() ListPersonResponse
}

type useCase struct {
	repo repository.PersonRepository
}

func NewListPersonUseCase(repo repository.PersonRepository) ListPersonUseCase {
	return &useCase{repo}
}

func (u *useCase) ListPerson() ListPersonResponse {
	list := u.repo.All()
	return &listPersonResponse{
		"Person list fetched",
		list,
	}
}
