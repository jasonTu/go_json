package main

import (
	"fmt"
	"flag"
	"go_json/gopl/ch7/bytecounter"
	"go_json/gopl/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 10.0, "the temprature")

func main() {
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)    // 5

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)    // 12

	// tempconv
    flag.Parse()
	fmt.Println(*temp)
}
