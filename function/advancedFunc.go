package main

import "fmt"
import "errors"

type operator func(x, y int) int
type calculator func(x, y int) (int, error)

func calculate(x int, y int, op operator) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

func genCalculate(op operator) calculator {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	op := func(x, y int) int {
		return x + y
	}
	x, y := 56, 78
	add := genCalculate(op)
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
