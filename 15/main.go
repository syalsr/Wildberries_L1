package main

import (
	"fmt"
)

/*
TODO К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить? Приведите корректный пример реализации.
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

/*
1. GC очистит данные глобальной переменной после выхода из программы
2. Если мы присваиваем что-то в глобальную переменную значит мы хотим
воспользоваться результатом где-то еще, а значит лучще возвращать строку
3. Если мы ничего больше не делаем с переменной v, то можем сразу вернуть
4. Если у нас символы строки кодируется больше чем одним байтом, то мы срежем 100 байт, а не
100 символов, как хотели
*/

func slice(str string, count int) (res []rune) {
	res = []rune(str)[:count]
	return
}

func yetAnotherSomeFunc() string {
	v := createHugeString(1 << 10)
	return string(slice(v, 10))
}
func main() {
	str := yetAnotherSomeFunc()
	fmt.Println(str)
}

func createHugeString(i int64) (s string) {
	for i > 0 {
		s += "и"
		i--
	}
	return
}
