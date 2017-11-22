package greeting

import "testing"

func TestGetMessage(t *testing.T) {
	expected := "Hello"
	resp := &greetingResponse{expected}

	if resp.GetGreeting() != expected {
		t.Errorf("Expected message %s, found %s", expected, resp.GetGreeting())
	}
}
