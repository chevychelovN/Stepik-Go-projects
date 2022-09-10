// Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
package main

import "fmt"

func main() {
    var n int
    var c int 
    var d int
    var i = 1
    fmt.Scan(&n)
    fmt.Scan(&c)
    fmt.Scan(&d)
    for ; i <= n; i++{ 
        if i % c == 0 && i % d != 0 {
            fmt.Print(i)
            break
        }
    }

}
