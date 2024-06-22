package chain

// BaseHandler provides a default implementation of the Handler interface.
type BaseHandler struct {
	next Handler
}

// SetNext sets the next handler in the chain.
func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

// Handle passes the request to the next handler in the chain if available.
func (b *BaseHandler) Handle(request string) string {
	if b.next != nil {
		return b.next.Handle(request)
	}
	return ""
}
