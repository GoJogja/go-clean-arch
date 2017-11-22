package greeting

import (
	"fmt"

	"github.com/arseto/go-clean-arch/domain/model"
)

type GreetingUseCase interface {
	GetGreeting(req GreetingRequest) GreetingResponse
}

type GreetingRequest interface {
	GetName() string
}

type greetingUseCase struct{}

func NewGreetingUseCase() GreetingUseCase {
	return &greetingUseCase{}
}

func (h *greetingUseCase) GetGreeting(req GreetingRequest) GreetingResponse {
	person := model.NewPerson(req.GetName())
	message := fmt.Sprintf(
		"Hello %s!",
		person.GetName(),
	)
	return &greetingResponse{message}
}
