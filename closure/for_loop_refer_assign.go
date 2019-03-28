package main

import "fmt"

func main() {
  var functions []func()

  for i := 0; i < 10; i++ {
    functions = append(functions, build(i))
  }

  for _, f := range functions {
    f()
  }
}

func build(val int) func() {
  return func() {
    fmt.Println(val)
  }
}

/*
package main

import "fmt"

func main() {
  var functions []func()

  for i := 0; i < 10; i++ {
    functions = append(functions, func() {
      fmt.Println(i)
    })
  }

  for _, f := range functions {
    f()
  }
}
*/
