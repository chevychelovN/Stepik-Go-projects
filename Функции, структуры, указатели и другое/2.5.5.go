// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

package main

import (
    "fmt"
    "unicode/utf8"
    "strings"
    "os"
    "bufio"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    rs := []rune(text)
    var counter int = 0
    for i:=0; i < utf8.RuneCountInString(string(rs)); i++ {
        counter = strings.Count(text, string(rs[i]))
        if counter > 1 {
            text = strings.ReplaceAll(text, string(rs[i]), "")
        }
    }
    fmt.Println(text)
}
