/*
Внутри функции main (объявлять функцию не нужно) необходимо написать программу:

На стандартный ввод подается 10 целых чисел, разделенных пробелами (числа могут повторяться). Для чтения из стандартного ввода импортирован пакет fmt.

Вам необходимо вычислить результат выполнения функции work для каждого из полученных чисел. Функция work имеет следующий вид:

func work(x int) int
Результаты вычислений , разделенные пробелами, должны быть напечатаны в строку.

Однако работа функции work занимает слишком много времени. Выделенного вам времени выполнения не хватит на последовательную обработку каждого числа, поэтому необходимо реализовать кэширование уже готовых результатов и использовать их в работе.

После завершения работы программы результат выполнения будет дополнен информацией о соблюдении установленного лимита времени выполнения.
*/

var a int
m1 := make(map[int]int)
for i := 0; i < 10; i++ {
    fmt.Scan(&a)
    if value, inMap := m1[a]; inMap {
	fmt.Print(value, " ")
    } else {
    m1[a] = work(a)
    fmt.Print(m1[a], " ")
}
}
