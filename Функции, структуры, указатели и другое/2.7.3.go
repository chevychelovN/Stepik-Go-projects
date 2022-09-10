// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

package main

import (
	"fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    switch {
        case strings.Contains(text, "9"):
            fmt.Print("9")
        case strings.Contains(text, "8"):
            fmt.Print("8")
        case strings.Contains(text, "7"):
            fmt.Print("7")
        case strings.Contains(text, "6"):
            fmt.Print("6")
        case strings.Contains(text, "5"):
            fmt.Print("5")
        case strings.Contains(text, "4"):
            fmt.Print("4")
        case strings.Contains(text, "3"):
            fmt.Print("3")
        case strings.Contains(text, "2"):
            fmt.Print("2")
        case strings.Contains(text, "1"):
            fmt.Print("1")
        case strings.Contains(text, "0"):
            fmt.Print("0")
}
}
