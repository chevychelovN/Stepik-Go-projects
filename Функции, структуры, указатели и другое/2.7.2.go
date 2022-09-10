// Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

package main

import (
	"fmt"
    "os"
    "bufio"
    "unicode/utf8"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    bs := []rune(text)
    i := 0
    j:= utf8.RuneCountInString(text)
    for i = i; i < j - 1; i++ {
        fmt.Print(string(bs[i]), "*")
    }
    fmt.Print(string(bs[i]))

}
