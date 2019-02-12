package main

import (
	"fmt"
)

type Foo string

func main() {
	type X map[string]int
	var x X
	var y map[string]int
	x = y
	fmt.Println(x)
	var a Foo
	var b string
	b = "a string"
	// _ = b
	// a = b
	a = Foo(b)
	// a = "string"
	fmt.Println(a)
	a = "b string"
	fmt.Println(a)
}
