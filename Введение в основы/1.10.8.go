//Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
package main

import "fmt"

func main() {
    var x int
    fmt.Scan(&x)
    var y int
    fmt.Scan(&y)
    var num int
    var numX int
    var xOld int = x
    var yOld int = y
    var i int
    switch {
        case x / 10000 > 0:
        i = 10000
        case x / 1000 > 0:
        i = 1000
        case x / 100 > 0:
        i = 100
        case x / 10 > 0:
        i = 10
        case x / 1 > 0:
        i = 1
    }
    for xOld > 0 {
        numX = xOld / i
        xOld = xOld - numX * i
        for yOld > 0 {
            num = yOld % 10
            yOld = yOld / 10
            if numX == num {
                fmt.Print(num)
                fmt.Print(" ")
            }
        }
        yOld = y
        i = i / 10
    }
}
