package main

import "fmt"

func main() {  
    x := 1
    fmt.Println(x)     //prints 1
    {
        fmt.Println(x) //prints 1
        x := 2
        fmt.Println(x) //prints 2
    }
    fmt.Println(x)     //prints 1 (bad if you need 2)
}

/*
Analysis:
➜  practice git:(master) ✗ go tool vet -shadow shadow_var.go 
shadow_var.go:10: declaration of "x" shadows declaration at shadow_var.go:6
*/
