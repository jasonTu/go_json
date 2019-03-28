package main

import "fmt"

func main() {
	var nestedFunctorInMain func()
	counter := 0
	nestedFunctorInMain = func() {
		counter += 1
		fmt.Println(counter)
	}
	nestedFunctorInMain()
	nestedFunctorInMain()
	nestedFunctorInMain()
}
