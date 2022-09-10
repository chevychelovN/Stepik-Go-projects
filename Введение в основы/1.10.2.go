//Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
//Вывести сумму всех чисел от A до B  включительно.
package main

import "fmt"

func main() {
    var a int
    var b int
    var c int = 0
    fmt.Scan(&a)
    fmt.Scan(&b)
    for a <= b {
        c = c + a
        a++
    }
    fmt.Println(c)

}
