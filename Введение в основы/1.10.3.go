//Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
package main

import "fmt"

func main() {
    var i int
    fmt.Scan(&i)
    var a int = 0
    var sum int = 0
    var num int = 0
for ; a < i; a++{
    fmt.Scan(&num)
    if num >= 10 && num <= 99 && num % 8 == 0 {
        sum = sum + num
    }
}
    fmt.Println(sum)
}
