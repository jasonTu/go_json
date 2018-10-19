package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Enter function main.")
	defer func(){
		fmt.Println("The first defer function.")
	}()
	defer func(){
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	defer func(){
		fmt.Println("The third defer function.")
	}()
	panic(errors.New("something wrong"))
	fmt.Println("Exit function main.")
}

/*
Result:
➜  panic git:(master) ✗ go run panicRecover.go
Enter function main.
The third defer function.
Enter defer function.
panic: something wrong
Exit defer function.
The first defer function.
 */
