package main

/*
import "fmt"

type T1 struct {
    name string
}
type T2 struct {
    name string
}
func main() {
    vs := []interface{}{T2(T1{"foo"}), string(322), []byte("abł")}
    for _, v := range vs {
        fmt.Printf("%v %T\n", v, v)
    }
}
*/

/*
{foo} main.T2
ł string
[97 98 197 130] []uint8}
 */

/*
import "fmt"

type T struct {
    name string
}

func main() {
    v1 := struct{ name string }{"foo"}
    fmt.Printf("%T\n", v1) // struct { name string }
    var v2 T
    v2 = v1
    fmt.Printf("%T\n", v2) // main.T
}
*/

/*
Output:
struct { name string }
main.T
 */

/*
type I1 interface {
    M1()
}
type I2 interface {
    M1()
}
type T struct{}
func (T) M1() {}
func main() {
    var v1 I1 = T{}
    var v2 I2 = v1
    _ = v2
}
*/

/*
type I1 interface {
    M1()
    M2()
}
type I2 interface {
    M1()
    I3
}
type I3 interface {
    M2()
}
type T struct{}
func (T) M1() {}
func (T) M2() {}
func main() {
    var v1 I1 = T{}
    var v2 I2 = v1
    _ = v2
}
*/

/*
type I1 interface {
    M1()
    M2()
}
type I2 interface {
    M1()
}
type T struct{}
func (T) M1() {}
func (T) M2() {}
func main() {
    var v1 I1 = T{}
    var v2 I2 = v1
    _ = v2
}
*/

/*
type I1 interface {
    M1()
}
type I2 interface {
    M1()
    M2()
}
type T struct{}
func (T) M1() {}
func main() {
    var v1 I1 = T{}
    var v2 I2 = v1
    _ = v2
}
*/

/*
Output:
# command-line-arguments
./interface_2.go:115:9: cannot use v1 (type I1) as type I2 in assignment:
        I1 does not implement I2 (missing M2 method)
*/

/*
type I1 interface {
    M1()
}
type T struct{}
func (T) M1() {}
func main() {
    var v1 I1 = T{}
    _ = v1
}
*/

/*
type I1 interface {
    M1()
}
type T struct{}
func (T) M1() {}
func main() {
    var v1 I1 = T{}
    var v2 T = v1
    _ = v2
}
*/

/*
Output:
# command-line-arguments
./interface_2.go:147:9: cannot use v1 (type I1) as type T in assignment: need type assertion
*/

/*
import "fmt"

type I interface {
    M()
}
type T struct {}
func (T) M() {}
func main() {
    var v I = T{}
    fmt.Println(T(v))
}
*/

/*
Output:
# command-line-arguments
./interface_2.go:168:18: cannot convert v (type I) to type T: need type assertion
*/

/*
import "fmt"

type I1 interface {
    M()
}
type I2 interface {
    M()
    N()
}
func main() {
    var v I1
    fmt.Println(I2(v))
}
*/

/*
Output:
# command-line-arguments
./interface_2.go:190:19: cannot convert v (type I1) to type I2:
        I1 does not implement I2 (missing N method)
*/

/*
import "fmt"

type I interface {
    M()
}
type T struct{}
func (T) M() {}
func main() {
    var v1 I = T{}
    v2 := v1.(T)
    fmt.Printf("%T\n", v2) // main.T
}
/*

/*
Output:
main.T
*/

/*
import "fmt"

type I interface {
    M()
}
type T1 struct{}
func (T1) M() {}
type T2 struct{}
func main() {
    var v1 I = T1{}
    v2 := v1.(T2)
    fmt.Printf("%T\n", v2)
}
*/

/*
Output:
# command-line-arguments
./interface_2.go:232:13: impossible type assertion:
        T2 does not implement I (missing M method)
*/

/*
import "fmt"

type I interface {
    M()
}
type T1 struct{}
func (T1) M() {}
type T2 struct{}
func (T2) M() {}
func main() {
    var v1 I = T1{}
    v2 := v1.(T2)
    fmt.Printf("%T\n", v2)
}
*/

/*
Output:
panic: interface conversion: main.I is main.T1, not main.T2

goroutine 1 [running]:
main.main()
        /root/go/src/go_json/interface/interface_2.go:256 +0x45
		exit status 2
*/

/*
import "fmt"

type I interface {
    M()
}
type T1 struct{}
func (T1) M() {}
type T2 struct{}
func (T2) M() {}
func main() {
    var v1 I = T1{}
    v2, ok := v1.(T2)
    if !ok {
        fmt.Printf("ok: %v\n", ok) // ok: false
        fmt.Printf("%v,  %T\n", v2, v2) // {},  main.T2
    }
}
*/

/*
Output:
ok: false
{},  main.T2}
*/

/*
type I1 interface {
    M()
}
type I2 interface {
    I1
    N()
}
type T struct{
    name string
}
func (T) M() {}
func (T) N() {}
func main() {
    var v1 I1 = T{"foo"}
    var v2 I2
    v2, ok := v1.(I2)
    fmt.Printf("%T %v %v\n", v2, v2, ok) // main.T {foo} true
}
*/

import "fmt"

type I1 interface {
    M1()
}
type T1 struct{}
func (T1) M1() {}
type I2 interface {
    I1
    M2()
}
type T2 struct{}
func (T2) M1() {}
func (T2) M2() {}
func main() {
    var v I1
    switch v.(type) {
    case T1:
            fmt.Println("T1")
    case T2:
            fmt.Println("T2")
    case nil:
            fmt.Println("nil")
    default:
            fmt.Println("default")
    }
}
