package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPointer := pointer.Interface().(*float64)
	converValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(converValue)
}

/*
(py2) ➜  reflect_sample git:(master) ✗ go run convert.go
0xc420088010
1.2345)
 */
