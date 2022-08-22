package main

/*
TODO Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}

	left, right := 0, len(arr)-1

	for left < right {
		midElem := (left + right) / 2
		if arr[midElem] == target {
			return arr[midElem]
		}
		if arr[midElem] > target {
			right = midElem - 1
		} else {
			left = midElem + 1
		}
	}

	return -1
}

func main() {

}
