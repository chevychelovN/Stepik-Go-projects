//Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
package main

import "fmt"

func main() {
    var num int
    var a int = 0
    var sum int = 0
    fmt.Scan(&num)
    for num != 0{
        if num > a {
            a = num
            sum = 0
        }
        if num == a {
            sum ++
        }
        fmt.Scan(&num)
}
    fmt.Println(sum)
}
