// Найдите самое большее число на отрезке от a до b, кратное 7 .

package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    var b int
    fmt.Scan(&b)
    var i int
    for i = b; i > a; i-- {
        if i % 7 == 0 {
            fmt.Print(i)
            break
        }
    }
    if i == a {
        fmt.Print("NO")
    }
}
