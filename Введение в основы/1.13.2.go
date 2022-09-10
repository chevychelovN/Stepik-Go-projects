// Дано трехзначное число. Переверните его, а затем выведите. 

package main
import "fmt"

func main() {
    var a int 
    fmt.Scan(&a)
    var array [3]int
    array[0] = a % 10
    a = a / 10
    array[1] = a % 10
    a = a / 10
    array[2] = a
    
    fmt.Print(array[0])
    fmt.Print(array[1])  
    fmt.Print(array[2])
}
