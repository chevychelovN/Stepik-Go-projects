// Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника. Нужно проверить, является ли треугольник прямоугольным. Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"

package main
import "fmt"

func main() {
    var a int 
    fmt.Scan(&a)
    var b int
    fmt.Scan(&b)
    var c int
    fmt.Scan(&c)
    if (a * a + b * b) == (c * c){
    fmt.Print("Прямоугольный")
    } else {
        fmt.Print("Непрямоугольный")
    }
}
