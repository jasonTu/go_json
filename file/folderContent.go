package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("test")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(files))
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
