package greeting

type GreetingResponse interface {
	GetGreeting() string
}

type greetingResponse struct {
	message string
}

func (r *greetingResponse) GetGreeting() string {
	return r.message
}
