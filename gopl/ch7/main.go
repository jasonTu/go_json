package main

import (
	"io"
	"os"
	"fmt"
	"flag"
	"bytes"
	"go_json/gopl/ch7/bytecounter"
	"go_json/gopl/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 10.0, "the temprature")
const debug = true

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

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

	// Report dynamic type of an interface value for debugging
	var w io.Writer
	fmt.Printf("%T\n", w)    // <nil>

	w = os.Stdout
	fmt.Printf("%T\n", w)    // *os.File

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)    // *bytes.Buffer

	// An interface Containing a NIL Pointer Is Non-NIL
	var buf *bytes.Buffer    // debug set to "false" will cause panic
	var buf io.Writer    // debug set to "false" will OK
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		// use buf
	}
	fmt.Printf("%s", buf)
}
