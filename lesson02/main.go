package main

import "fmt"

func main() {
	ch := make(chan string,2)
	defer close(ch)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}