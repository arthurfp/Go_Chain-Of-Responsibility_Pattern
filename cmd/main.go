package main

import (
	"chain-of-responsibility-pattern-go/pkg/chain"
	"fmt"
)

func main() {
	handlerA := &chain.ConcreteHandlerA{}
	handlerB := &chain.ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	fmt.Println(handlerA.Handle("A"))
	fmt.Println(handlerA.Handle("B"))
	fmt.Println(handlerA.Handle("C"))
}
