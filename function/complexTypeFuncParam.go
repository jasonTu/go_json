package main

import "fmt"

func modifyArray(arr [3][]string) [3][]string {
	arr[1] = []string{"x", "x", "x"}
	return arr
}

func modifyArray2(arr [3][]string) [3][]string {
	arr[1][1] = "x"
	return arr
}

func main() {
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	fmt.Printf("The array: %v\n", complexArray1)
	// array2 := modifyArray(complexArray1)
	array2 := modifyArray2(complexArray1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", complexArray1)
}
