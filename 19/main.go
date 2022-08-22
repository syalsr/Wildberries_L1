package main

/*
TODO Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode
*/

func reverseStr(str string) (result string) {
	for idx, _ := range str {
		result += string(str[len(str)-idx-1])
	}
	return
}
func yetAnotherReverseStr(str string) string {
	nstr := []rune(str)
	for idx := 0; idx < len(str)/2; idx++ {
		nstr[idx], nstr[len(nstr)-idx-1] = nstr[len(nstr)-idx-1], nstr[idx]
	}
	return string(nstr)
}
func main() {

}
