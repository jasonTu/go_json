package main

import "fmt"

func modifyArray(arr [3]string) [3]string {
	arr[1] = "x"
	return arr
}

func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
}
