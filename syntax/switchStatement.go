package main

import "fmt"

func main() {
	/*
	// "1 + 3" type differ from value1[index], bellow code can't compile ok
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 3 {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4, 5 or 6")
	}
	*/

	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4, 5 or 6")
	}

	/*
	// 不能编译成功，因为在case语句中存在结果相等的子表达式
	value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value3[4] {
	case 0, 1, 2:
		fmt.Println("0 or 1 or 2")
	case 2, 3, 4:
		fmt.Println("2 or 3 or 4")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
	*/

	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or 6")
	}
}
