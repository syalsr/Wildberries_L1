package main

import "fmt"

/*
TODO Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

func main() {
	stdout := make(chan int)
	go func() {
		arr := []int{2, 4, 6, 8, 10}
		for _, elem := range arr {
			stdout <- elem * elem
		}
		close(stdout)
	}()

	var sumSquaring int
	for elem := range stdout {
		sumSquaring += elem
	}
	fmt.Print(sumSquaring)

}
