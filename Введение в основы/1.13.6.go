// Даны два числа. Найти их среднее арифметическое.

package main
import "fmt"

func main() {
    var a float64 
    fmt.Scan(&a)
    var b float64
    fmt.Scan(&b)
    var c float64 = (a + b) / 2
    fmt.Print(c)
}
