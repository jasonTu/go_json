package main

import (
	"fmt"
	"strings"
)

func main() {
	// str := "10.1.2.23, 10.1.2.23, 114.179.206.171"
	str := "10.1.2.23"
	ips := strings.Split(str, ",")
	for _, ip := range ips {
		fmt.Println(ip)
		fmt.Println(strings.TrimSpace(ip))
	}
}
