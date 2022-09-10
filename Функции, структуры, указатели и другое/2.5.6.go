// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

package main

import (
    "fmt"
    "unicode/utf8"
    "unicode"
    "os"
    "bufio"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    rs := []rune(text)
    if utf8.RuneCountInString(string(rs)) < 5 {
        fmt.Print("Wrong password")
        return
    }
    for i:=0; i < utf8.RuneCountInString(string(rs)); i++ {
        if unicode.IsDigit((rs[i])) == false && unicode.Is(unicode.Latin, rs[i]) == false {
            fmt.Print("Wrong password")
            return
        }
    }
    fmt.Print("Ok")
}
