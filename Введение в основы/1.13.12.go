// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

package main

import "fmt"
import "math"

func main() {
    var a float64
    fmt.Scan(&a)
    var b float64 = 2
    var i float64 
    for i = 0; math.Pow(b, i) <= a; i++ {
        fmt.Print(math.Pow(b, i), " ")
    }
    
}
