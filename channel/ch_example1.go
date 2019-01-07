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
	for num := range ch {
		fmt.Printf("F(%d): \t%d\n", idx, num)
		idx++
	}
}

/*
Output:
F(0):   0
F(1):   1
F(2):   1
F(3):   2
F(4):   3
F(5):   5
F(6):   8
F(7):   13
F(8):   21
F(9):   34
F(10):  55)
*/
