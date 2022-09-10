// Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

package main

import "fmt"


func main() {
    var a int
    fmt.Scan(&a)
    var b int = 0
    var c int = 1
    var d int = 0
    var num int = 0
    for b <= a {
                if b == a {
            fmt.Print(num)
                    break
        }
        d = b + c
        b = c
        c = d
        num++
    }
    if b > a {
        fmt.Print("-1")
    }
}
