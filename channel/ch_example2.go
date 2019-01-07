/*
Refer:
Range over channels in Ga: https://medium.com/golangspec/range-over-channels-in-go-3d45dd825d2a
*/
package main

import "fmt"

func FibonacciProducer(ch chan int, count int) {
	n2, n1 := 0, 1

	for count >= 0 {
		ch <- n2
		count--
		n2, n1 = n1, n2+n1
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go FibonacciProducer(ch, 10)
	idx := 0
	// Index variable is not allowed when range expression is a channel 
	for idx, num := range ch {
		fmt.Printf("F(%d): \t%d\n", idx, num)
		idx++
	}
}

/*
Output:
# command-line-arguments
./ch_example2.go:21:18: too many variables in range
*/
