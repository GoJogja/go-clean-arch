package greeting

import "testing"

type mockGreetingRequest struct{}

func (m *mockGreetingRequest) GetName() string {
	return "Arseto"
}

func TestGetGreeting(t *testing.T) {

	uc := NewGreetingUseCase()
	req := &mockGreetingRequest{}

	resp := uc.GetGreeting(req)

	expected := "Hello Arseto!"

	if resp.GetGreeting() != expected {
		t.Errorf("Expected message %s, found %s", expected, resp.GetGreeting())
	}

}
