// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

func minimumFromFour() int {
    var min int
    var num int
    fmt.Scan(&min)
    for i := 0; i < 3; i++ {
        fmt.Scan(&num)
        if num < min {
            min = num
        }
    }
    return min
}
