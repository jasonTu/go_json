package main

import "log"
import "time"

func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	time.Sleep(1 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) int {
	return x + x
}

func doubleEx(x int) (result int) {
	defer func() {
		log.Printf("double(%d) = %d", x, result)
	}()

	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += x
		log.Printf("triple(%d) = %d", x, result)
	}()
	return double(x)
}

func main() {
	bigSlowOperation()
	_ = doubleEx(4)
	_ = triple(4)
}

/*
Output:
➜  ch5 git:(master) ✗ go run defer_func.go
2018/12/20 03:00:55 enter bigSlowOperation
2018/12/20 03:00:56 exit bigSlowOperation (1.000423885s)
2018/12/20 03:00:56 double(4) = 8
2018/12/20 03:00:56 triple(4) = 12
*/
