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

func (u User) ReflectCallFunc() {
	fmt.Println("In ReflectCallFunc of User")
}

func main() {
	user := User{1, "Jason Tu", 25}

	DoFieldAndMethod(user)
}

func DoFieldAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get type is: ", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get value is: ", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

/*
Result:
(py2) ➜  reflect_sample git:(master) ✗ go run detect.go
get type is:  User
get value is:  {1 Jason Tu 25}
Id: int = 1
Name: string = Jason Tu
Age: int = 25
ReflectCallFunc: func(main.User))
 */
