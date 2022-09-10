// Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.

func vote(x int, y int, z int) int {
    var a int = x
    var b int = y
    if a == b {
        return a
    }
    return z
    
}
