package greeting

type greetingRequest struct {
	Name string
}

func (h *greetingRequest) GetName() string {
	return h.Name
}
