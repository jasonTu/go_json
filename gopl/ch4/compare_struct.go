package main

import "fmt"

type address struct {
	hostname    string
	port        int
}

func main() {
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}
