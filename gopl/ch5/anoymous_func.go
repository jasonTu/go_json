package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x*x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

/*
Output:
➜  ch5 git:(master) ✗ go run anoymous_func.go
1
4
9
 */
