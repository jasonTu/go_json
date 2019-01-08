package main

import (
	"fmt"
	"sync"
)

func main() {
	var v int
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go func() {
		v = 1
		<-ch
		wg.Done()
	}()
	go func() {
		ch <- 1
		ch <- 1
		ch <- 1
		ch <- 1
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}
