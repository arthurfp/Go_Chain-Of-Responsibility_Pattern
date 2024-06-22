package main

import (
	"chain-of-responsibility-pattern-go/pkg/chain"
	"fmt"
)

func main() {
	baseHandler := &chain.BaseHandler{}
	fmt.Println(baseHandler.Handle("Test Request"))
}
