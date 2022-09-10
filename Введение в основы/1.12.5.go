//Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.

package main
import "fmt"

func main() {
    var a int 
    fmt.Scan(&a)
    slice := make([]int, a, a)
    for i:=0; i < a; i++ {
        fmt.Scan(&slice[i])
    }
    var even int = 0
        for i:=0; i < a; i++ {
            if slice[i] > 0 {
                even++
            }
    }
    fmt.Print(even) 
    
}
