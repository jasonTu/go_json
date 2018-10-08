package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345
	fmt.Println("Old value of pointer:", num)

	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer:", newValue.Type)
	fmt.Println("settablity of pointer:", newValue.CanSet())

	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

	pointer = reflect.ValueOf(num)
	// newValue = pointer.Elem() // this will panic when the pointer is not a
	// real pointer
}

/*
Result:
(py2) ➜  reflect_sample git:(master) ✗ go run update.go
Old value of pointer: 1.2345
type of pointer: 0x4825f0
settablity of pointer: true
new value of pointer: 77
 */
