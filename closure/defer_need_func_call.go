package main

import "fmt"

func main() {
	defer setup()()
}

func setup() func() {
  fmt.Println("pretend to set things up")

  return func() {
    fmt.Println("pretend to tear things down")
  }
}

/*
func main() {
  defer func() {
    fmt.Prinltn("teardown")
  }
}
*/
