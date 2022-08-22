package main

/*
TODO Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/

func createSet(arr []string) map[string]int { //O(N)
	m := make(map[string]int)
	for _, elem := range arr { //O(N)
		m[elem] += 1
	}
	return m
}
func createSet1(arr []string) (res []string) { //O(NlogN)
	if len(arr) <= 1 {
		return arr
	}
	quickSort(arr) //O(NlogN)
	pos := 0
	for _, item := range arr { //O(N)
		if res[pos] != item {
			res = append(res, item)
			pos++
		}
	}
	return
}

func main() {

}
