package chain

// Handler defines the interface for handling requests.
type Handler interface {
	SetNext(handler Handler)
	Handle(request string) string
}
