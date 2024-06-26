package chain

import "testing"

func TestConcreteHandlerA_Handle(t *testing.T) {
	handler := &ConcreteHandlerA{}
	baseHandler := &BaseHandler{}
	handler.SetNext(baseHandler)

	result := handler.Handle("A")
	expected := "ConcreteHandlerA: Handled A"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

	result = handler.Handle("B")
	expected = ""

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
