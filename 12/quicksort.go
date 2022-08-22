package main

import (
	"math/rand"
	"time"
)

func quickSort(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	pivot := &arr[r.Intn(len(arr))]
	capacity := (len(arr) / 2) - 1
	lhsArr := make([]string, 0, capacity) //все что меньше пивота
	rhsArr := make([]string, 0, capacity) //все что больше пивота
	//middle := make([]int, 0, cap) //все что равно пивоту

	for idx, item := range arr {
		if &arr[idx] == pivot {
			continue
			//middle = append(middle, item)
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
