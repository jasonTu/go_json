package main

var initArg [20]int

func init() {
	initArg[0] = 10
	for i := 1; i < len(initArg); i++ {
		initArg[i] = initArg[i-1] * 2
	}
}

func main() {
	_ = initArg[0]
}
