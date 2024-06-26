package chain

import "testing"

func TestConcreteHandlerB_Handle(t *testing.T) {
	handler := &ConcreteHandlerB{}
	baseHandler := &BaseHandler{}
	handler.SetNext(baseHandler)

	result := handler.Handle("B")
	expected := "ConcreteHandlerB: Handled B"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	result = handler.Handle("A")
	expected = ""

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
