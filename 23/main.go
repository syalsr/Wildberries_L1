package main

/*
TODO Удалить i-ый элемент из слайса.
*/

func deleteElemFromSlice(arr []int64, idx uint) []int64 {
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr
}

func main() {

}
