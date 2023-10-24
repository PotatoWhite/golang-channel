package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	for {
		fmt.Println("Try some text and press enter. Press 'q' to quit.")
		fmt.Print("> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == "q" {
			break
		}

		ping <- userInput

		response := <-pong
		fmt.Printf("From pong: %s\n", response)
	}

	close(ping)
	close(pong)

	fmt.Println("Bye!")
}
