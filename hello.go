package main

import (
	"fmt"
	"github.com/matthewfinger/golang_greetings"
)

func main() {
	message := golang_greetings.Hello("Matthew")
	fmt.Println(message)
	fmt.Println(golang_greetings.Greet("Evan", false))
}