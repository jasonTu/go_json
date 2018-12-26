package main

import (
	"os"
	"fmt"
	"errors"
	. "go_json/gopl/ch7/errortype"
)

func main() {
	var err1 error
	perror := PathError{"open", "/no/such/file", errors.New("No such file or directory")}
	err1 = perror
	fmt.Println(perror)    // open /no/such/file No such file or directory
	fmt.Println(err1)    // open /no/such/file No such file or directory

	_, err := os.Open("/no/such/file")
	fmt.Println(err)    // open /no/such/file: no such file or directory

	// nerr := perror.(error)
	// fmt.Println(nerr)
}
