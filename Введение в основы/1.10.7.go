//Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)
    var p int
    fmt.Scan(&p)
    var y int
    fmt.Scan(&y)
    var i int = 0
    for x = x; x < y; x = x + p * x/100 { 
        i++
    }
    fmt.Print(i)
}
