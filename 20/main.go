package main

import "strings"

/*
TODO Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverse(str string) (result string) {
	arrStr := strings.Split(str, " ")

	for i := 0; i < len(arrStr)/2; i++ {
		arrStr[i], arrStr[len(arrStr)-i-1] = arrStr[len(arrStr)-i-1], arrStr[i]
	}
	return strings.Join(arrStr, " ")
}

func main() {

}
