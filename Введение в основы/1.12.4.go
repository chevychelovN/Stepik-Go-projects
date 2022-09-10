//Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

package main
import "fmt"

func main() {
    var a int 
    fmt.Scan(&a)
    slice := make([]int, a, a)
    for i:=0; i < a; i++ {
        fmt.Scan(&slice[i])
    }
        for i:=0; i < a; i = i + 2 {
        fmt.Print(slice[i], " ")
    }
        
    
}
