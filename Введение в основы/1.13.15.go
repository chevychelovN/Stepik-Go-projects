// Из натурального числа удалить заданную цифру.

package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    var num int
    fmt.Scan(&num)
    var b []int
    for a > 0 {
        b = append(b, a%10)
        a = a / 10
    }
    var c []int
    for i := len(b)-1; i >= 0; i-- {
        if b[i] != num {
            c = append(c, b[i])
        }
    }
    for i := 0; i < len(c); i++ {
        fmt.Print(c[i])
    }
}
