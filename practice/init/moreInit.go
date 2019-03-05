package main

import (
	"fmt"
)

func init() {
	fmt.Println("a init")
}

func init() {
	fmt.Println("b init")
}

func init() {
	fmt.Println("c init")
}

func main() {
	fmt.Println("In main")
}
