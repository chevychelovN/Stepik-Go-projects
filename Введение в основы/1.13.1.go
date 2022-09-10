//Дано трехзначное число. Найдите сумму его цифр. 
package main
import "fmt"

func main() {
    var a int 
    fmt.Scan(&a)
    var sum int
    sum = sum + a % 10
    a = a / 10
        sum = sum + a % 10
   a = a / 10
        sum = sum + a % 10
    a = a / 10
    fmt.Print(sum)
    
}
