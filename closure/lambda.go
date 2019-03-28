package main

import "fmt"

func main() {
	lambda := func() {
		fmt.Println("I'm an anonymous function")
	}
	lambda()
}
