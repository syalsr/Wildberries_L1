package main

import "strings"

/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/

func checkUnqueString(str string) bool {
	m := make(map[string]int)
	str = strings.ToLower(str)
	for _, ch := range str {
		if _, inMap := m[string(ch)]; inMap {
			return false
		} else {
			m[string(ch)] = 1
		}
	}
	return true
}

func main() {

}
