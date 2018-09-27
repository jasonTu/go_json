package main

import "fmt"

func main() {
	number1 := []int{1, 2, 3, 4, 5, 6}
	for i := range number1 {
		if i == 3 {
			number1[i] |= i
		}
	}
	fmt.Println(number1)

	number2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(number2) - 1
	for i, e := range number2 {
		if i == maxIndex2 {
			number2[0] += e
		} else {
			number2[i+1] += e
		}
	}
	fmt.Println(number2)
}
