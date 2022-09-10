// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число. 

//Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

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
    for i = i; i < j; i++ {
        fmt.Print((int(bs[i]) - 48) * (int(bs[i]) - 48))
    }

}
