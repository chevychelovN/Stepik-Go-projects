// На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

// Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

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
    flag := 0
    size:= utf8.RuneCountInString(text)
    j:= utf8.RuneCountInString(text)
    for i = i; i < j; {
        for j = j - 1; j > i; j --{
            if bs[i] != bs[j] {
                fmt.Print("Нет")
                flag = 1
                break
            }
            i++
        }
        if flag == 1 {
            break
        }
    }
    if i == j && size % 2 == 1{
        fmt.Print("Палиндром")
    }
    if i == j + 1 && size % 2 == 0 {
        fmt.Print("Палиндром")
    }

}
