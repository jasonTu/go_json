package main

import (
	"fmt"
	"sync"
)


func main() {
	var v int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		v = 1
		wg.Done()
	}()
	go func() {
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}

/*
Output:
0 or 1
*/
