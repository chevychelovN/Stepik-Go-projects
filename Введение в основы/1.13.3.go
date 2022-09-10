// Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток. Например, если

//k=13257=3*3600+40*60+57,

//то h=3 и m=40.

package main
import "fmt"

func main() {
    var a int 
    fmt.Scan(&a)
    var h int = a / 3600
    a = a - h * 3600
    var m int = a / 60
    fmt.Print("It is ", h, " hours ", m, " minutes.")
}
