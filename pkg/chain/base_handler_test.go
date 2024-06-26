package chain

import "testing"

func TestBaseHandler_Handle(t *testing.T) {
	handler := &BaseHandler{}

	result := handler.Handle("Test")
	expected := ""

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
