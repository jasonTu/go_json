package main

import (
	"fmt"
)

var aInt int = 10
var aSliceInt [12]int

func init() {
	fmt.Println("Execute point init")
	for i := 0; i < len(aSliceInt); i++ {
		aSliceInt[i] = i
	}
}

func main() {
	fmt.Println("Execute point main")
	fmt.Println(aSliceInt)
	init()
}
