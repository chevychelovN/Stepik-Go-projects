// По данным числам, определите количество чисел, которые равны нулю.  

package main
import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    var b int = 0
    var num int
    for i := 0; i < a; i++ {
        fmt.Scan(&num)
        if num == 0 {
            b++
        }
    }
    fmt.Print(b)
}
