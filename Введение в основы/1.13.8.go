// Найдите количество минимальных элементов в последовательности.

package main

import "fmt"

func main() {
    var num int = 100000
    var b int
    fmt.Scan(&b)
    var sum int = 0
        var a int = num
    for i:= 0; i < b; i++{
                fmt.Scan(&num)
        if num < a {
            a = num
            sum = 0
        }
        if num == a {
            sum ++
        }
}
    fmt.Println(sum)
}
