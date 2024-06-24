package chain

// ConcreteHandlerA handles requests it is responsible for and passes others to the next handler.
type ConcreteHandlerA struct {
	BaseHandler
}

// Handle processes the request or passes it to the next handler.
func (h *ConcreteHandlerA) Handle(request string) string {
	if request == "A" {
		return "ConcreteHandlerA: Handled " + request
	}
	return h.BaseHandler.Handle(request)
}
