// На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

package main

import "fmt"
import "math"

func main() {
	var a float64
    var b float64
    fmt.Scan(&a)
    fmt.Scan(&b)
    var c float64 = math.Sqrt(a*a+b*b)
    fmt.Print(c)
}
