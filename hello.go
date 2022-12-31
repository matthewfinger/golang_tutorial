package main

import (
	"fmt"
	"github.com/matthewfinger/golang_greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	var name string
	fmt.Print("Enter a name: ")
	fmt.Scanf("%s", &name)
	message, err := golang_greetings.Hello(name)
	if (err != nil) {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", message)
	fmt.Println(golang_greetings.Greet("Evan", false))
}