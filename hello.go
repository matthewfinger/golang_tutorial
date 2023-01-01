package main

import (
	"fmt"
	"github.com/matthewfinger/golang_greetings"
	"log"
	"strings"
	"bufio"
	"os"
)


func NameLoop() {
	var input string
	fmt.Print("Enter a name (or multiple separated by commas): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println("You entered nothing or something stupid")
		NameLoop()
		return
	}

	input = scanner.Text()

	if (strings.ToLower(input) == "exit") {
		fmt.Println("Bye bye")
		return
	}
	components := strings.Split(input, ",")
	names := make([]string, 0, len(components))
	for _, component := range components {
		cleaned := strings.TrimSpace(component)
		if (cleaned != "") {
			names = append(names, cleaned)
		}
	}

	messages, err := golang_greetings.Hellos(names)
	if (err != nil) {
		fmt.Println("You entered nothing or something stupid")
		NameLoop()
		return
	}
	for _, message := range messages {
		fmt.Printf("%v\n", message)
	}
	if (input != "exit") {
		NameLoop()
	}
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	NameLoop()
}