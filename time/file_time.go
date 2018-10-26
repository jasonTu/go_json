package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	fh, _ := os.Open("file_time.go")
	defer fh.Close()

	fi, _ := fh.Stat()
	fmt.Printf("file modify time: %d\n", fi.ModTime().Unix())
	timestamp := time.Now().Unix()
	fmt.Printf("Current time is %d\n", timestamp)
}
