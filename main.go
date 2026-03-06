package main

import (
	"fmt"
	"hello/hello"
)

func main() {
	messages := hello.Hello("Lan") + " from Main"
	fmt.Println(messages)
}
