// По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    if a % 10 == 1 && (a > 19 || a < 10) {
        fmt.Print(a, " korova")
    }
    if a % 10 > 1 && a % 10 < 5 && (a > 19 || a < 10) {
        fmt.Print(a, " korovy")
    }
    if a % 10 > 4 && a % 10 < 10 || a % 10 == 0 || (a > 10 && a < 20){
        fmt.Print(a, " korov")
    }
}
