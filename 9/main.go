package main

import "fmt"

/*
TODO Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться
в stdout.
*/

func main() {
	arr := []int{3, 1, 2, 1, 2, 3, 4, 5, 10}
	chan1, chan2 := make(chan int), make(chan int)

	go func() {
		for _, item := range arr {
			chan1 <- item
			chan2 <- item * 2
		}
		close(chan1)
		close(chan2)
	}()
	go func() {
		for range chan1 { //выводим значения из первого канала, чтобы продолжить пушить туда значения
			<-chan1
		}
	}()
	for i := range chan2 {
		fmt.Println(i)
	}
}
