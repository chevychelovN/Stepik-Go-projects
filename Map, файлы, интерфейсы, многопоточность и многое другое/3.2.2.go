/*
Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.


Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше. Поэтому он решил подшутить над вами и подсунул вам свинью. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.


Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы, вам ничего импортировать не нужно!

Считывать и выводить ничего не нужно!

Ваша функция должна называться adding() !

Считайте что функция и пакет main уже объявлены!
*/

func adding(str, str1 string) int64 {
    rs := []rune(str)
    var res string
    rs1 := []rune(str1)
    var res1 string
    for i := range rs {
        if unicode.IsDigit(rs[i]) {
            res += string(rs[i]) 
        }
    }
        for i := range rs1 {
        if unicode.IsDigit(rs1[i]) {
            res1 += string(rs1[i])
        }
    }
    result1, _ := strconv.ParseInt(res1, 10, 64)
    result, _ := strconv.ParseInt(res, 10, 64)
    return result + result1
}
