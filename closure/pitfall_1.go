package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	defer func() {
		close(ch)
	}()
	for _, value := range slice {
		go func(v int) {
			ch <- v
		}(value)
	}
	for i := 0; i < len(slice); i++ {
		x := <- ch
		fmt.Println(x)
	}
}

/*
Result:
➜  closure git:(master) ✗ go run pitfall_1.go
6
2
1
3
5
4
 */
