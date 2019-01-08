package main

import (
	"fmt"
	"sync"
)


func main() {
	var v int
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		v = 1
		ch <- 1
		wg.Done()
	}()
	go func() {
		<-ch
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}

/*
The new thing is ch channel. Since receiving happens after
sending value to a channel and sending value happens after
assigning to v then above program prints always 1:
Output:
0 or 1
*/
