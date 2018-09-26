package main

import "fmt"
import "time"

func main() {
	intChan := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Printf("The candidate case is closed.\n")
			break
		}
		fmt.Printf("The candidate case is selected.\n")
	}
}
