package main

import "fmt"
import "time"

func main() {
	for i := 0; i < 10; i++ {
		/* 
		 * go语句后跟随的函数会被go routine调度器按顺序启用调度G，
		 * 但是执行顺序无法确定，而routine实际上是一个用户级的运行里程，
		 * 只有在执行时才会读取i当时的值，所以可能输出会乱序。
		*/
		go func() {
			fmt.Println(i)
		}()
	}
	// 如果不加下面这句，则大部分情况还没等到routines运行，主程序就已经退出了
	time.Sleep(time.Second)
}
