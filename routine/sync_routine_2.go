package main

import (
	"fmt"
	"sync"
)

func main() {
	var v, w int
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		v = 1
		ch <- 1
		fmt.Println(w)
		wg.Done()
	}()
	go func() {
		w = 2
		<-ch
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}
