package main

import "fmt"

func main() {
	var functor func() int
	functor = func() int {
		return 10
	}
	fmt.Println(functor())
	functor = func() int {
		return 15
	}
	fmt.Println(functor())
	functor = func() int {
		return 20
	}
	fmt.Println(functor())
}
