package main

import (
	"fmt"
	"os"
	"strconv"
	"go_json/gopl/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

/*
Build the binary anywhere: go build go_json/gopl/ch2/cf
➜  gopl git:(master) ✗ ./cf 32
32°F = 0°C, 32°C = 89.6°F
 */
