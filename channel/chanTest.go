package main

import "fmt"

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}
}
