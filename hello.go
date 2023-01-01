package main

import (
	"fmt"
	"github.com/matthewfinger/golang_greetings"
	"log"
	"strings"
)

func NameLoop() {
	var name string
	fmt.Print("Enter a name: ")
	fmt.Scanf("%s", &name)
	if (strings.ToLower(name) == "exit") {
		fmt.Println("Bye bye")
		return
	}
	message, err := golang_greetings.Hello(name)
	if (err != nil) {
		fmt.Println("You entered nothing or something stupid")
		NameLoop()
	}
	fmt.Printf("%v\n", message)
	if (name != "exit") {
		NameLoop()
	}
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	NameLoop()
}