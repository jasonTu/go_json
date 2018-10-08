package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345
	
	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))
}

/*
Result:
(py2) ➜  reflect_sample git:(master) ✗ go run basicFunc.go
type:  float64
value:  1.2345)
 */
