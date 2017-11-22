package greeting

import "testing"

func TestGetMessage(t *testing.T) {
	expected := "Arseto"
	req := &greetingRequest{expected}

	if req.GetName() != expected {
		t.Errorf("Expected message %s, found %s", expected, req.GetName())
	}
}
