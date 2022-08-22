package main

/*
TODO Реализовать пересечение двух неупорядоченных множеств.
*/
func Map(arr1 []int, arr2 []int) map[int]byte { //O(N) N - длина наибольшего массива
	m := make(map[int]byte)
	for _, item := range arr1 { //O(N) N - длина массива arr1
		m[item] += 1
	}
	for _, item := range arr2 { //N(M) M - длина массива arr2
		m[item] += 1
	}

	for key, value := range m { //O(k) K - количество элементов после пересечения
		if value == 1 {
			delete(m, key)
		}
	}
	return m
}

func Array(arr1 []int, arr2 []int) []int { //O(NlogN) N - длина наибольшего массива
	if len(arr1) == 0 || len(arr2) == 0 {
		return []int{}
	}

	arr1 = quickSort(arr1) //O(NlogN)
	arr2 = quickSort(arr2) //O(MlogM)

	arr3 := make([]int, 0, len(arr1)+len(arr2))

	pArr1, pArr2 := 0, 0
	for pArr1 < len(arr1) && pArr2 < len(arr2) { //O(N) N - длина наименьшего массива
		if arr1[pArr1] > arr2[pArr2] {
			if pArr2 != len(arr2) {
				pArr2++
			}
		} else if arr1[pArr1] < arr2[pArr2] {
			if pArr1 != len(arr1) {
				pArr1++
			}
		} else if arr1[pArr1] == arr2[pArr2] {
			arr3 = append(arr3, arr1[pArr1])
			if pArr1 != len(arr1) && pArr2 != len(arr2) {
				pArr1++
				pArr2++
			}
		}
	}

	return arr3
}
func main() {

}
