// Refer: http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#closure_for_it_vars
package main

import (  
    "fmt"
    "time"
)

/*
func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        go func() {
            fmt.Println(v)
        }()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: three, three, three
}
*/

/*
func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        go func() {
            fmt.Println(v)
        }()
		time.Sleep(1 * time.Second)
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}
*/

/*
func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        vcopy := v //
        go func() {
            fmt.Println(vcopy)
        }()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three in random order
}
*/

func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        go func(in string) {
            fmt.Println(in)
        }(v)
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three in random order
}
