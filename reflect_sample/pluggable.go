package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age: ", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func main() {
	user := User{1, "Jason Tu", 25}

	getValue := reflect.ValueOf(user)

	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("budeliao"), reflect.ValueOf(30)}
	methodValue.Call(args)
	// methodValue.Call("liaobude", 35) // want ([]reflect.Value)

	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}

/*
Result:
(py2) ➜  reflect_sample git:(master) ✗ go run pluggable.go
ReflectCallFuncHasArgs name:  budeliao , age:  30 and origal User.Name: Jason Tu
ReflectCallFuncNoArgs)
 */
