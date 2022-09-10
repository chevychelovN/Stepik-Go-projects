//  На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

package main

import (
    "fmt"
    "unicode/utf8"

)

func main() {
    var text string 
    fmt.Scan(&text)
    rs := []rune(text)
    for i:=1; i < utf8.RuneCountInString(text); i += 2 {
        fmt.Print(string(rs[i]))
    }
}
