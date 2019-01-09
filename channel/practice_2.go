package main

import (
	"fmt"
)

func main() {
	s := []int{5, 3, 1, 6, 0, 4}
	done := make(chan bool)
	// doSort is a lambda function, so a closure which knows the channel done
	doSort := func(s []int) {
		// sort(s)
		done <- true
	}

	// i := pivot(s)
	i := 2
	go doSort(s[:i])
	go doSort(s[i:])
	<-done
	<-done
	fmt.Println(s)
}
