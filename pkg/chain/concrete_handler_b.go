package chain

// ConcreteHandlerB handles requests it is responsible for and passes others to the next handler.
type ConcreteHandlerB struct {
	BaseHandler
}

// Handle processes the request or passes it to the next handler.
func (h *ConcreteHandlerB) Handle(request string) string {
	if request == "B" {
		return "ConcreteHandlerB: Handled " + request
	}
	return h.BaseHandler.Handle(request)
}
