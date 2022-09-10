// На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong

package main

import (
	"fmt"
    "os"
    "bufio"
    "strings"
    "unicode"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    bs := []rune(text)
    if  strings.HasSuffix(text, ".") && unicode.IsUpper(bs[0]) {
        fmt.Print("Right")
    } else {
        fmt.Print("Wrong")
    }
}
