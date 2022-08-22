package main

import (
	"math/rand"
	"time"
)

/*
TODO Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	pivot := &arr[r.Intn(len(arr))]

	capacity := (len(arr) / 2) - 1
	lhsArr := make([]int, 0, capacity) //все что меньше пивота
	rhsArr := make([]int, 0, capacity) //все что больше пивота

	for idx, item := range arr {
		if &arr[idx] == pivot {
			continue
		} else if item > *pivot {
			rhsArr = append(rhsArr, item)
		} else {
			lhsArr = append(lhsArr, item)
		}
	}

	lhsArr = quickSort(lhsArr)
	rhsArr = quickSort(rhsArr)

	lhsArr = append(lhsArr, *pivot)
	lhsArr = append(lhsArr, rhsArr...)

	return lhsArr
}

func main() {

}
