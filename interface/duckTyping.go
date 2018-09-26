package main

import "fmt"

type Dog struct {
	name string
	category string
}

type Pet interface {
	Name() string
	Category() string
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return dog.category
}

func main() {
	dog := Dog{"little dog", "dog category"}
	// var pet Pet = &dog
	var pet Pet = dog
	fmt.Printf("pet name: %s, category: %s.\n", pet.Name(), pet.Category())
	dog.name = "Dog King"
	fmt.Printf("pet name: %s, category: %s.\n", pet.Name(), pet.Category())
}
