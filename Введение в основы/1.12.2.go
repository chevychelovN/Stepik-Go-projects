//Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

package main
import "fmt"

func main() {
    var a int 
    fmt.Scan(&a)
    slice := make([]int, a, a)
    for i:=0; i < a; i++ {
        fmt.Scan(&slice[i])
    }
    fmt.Print(slice[3])
        
    
}
