package main

import (
	"fmt"
	"sync"
)

/*
TODO Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Inc struct {
	inc int64
	b   sync.Mutex
}

func (i *Inc) Add() {
	i.b.Lock()
	defer i.b.Unlock()
	i.inc++
}

func main() {
	incrementer := new(Inc)
	wg := &sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			ii := 0
			for ii < 1000 {
				incrementer.Add()
				ii++
			}
		}()
	}

	wg.Wait()

	fmt.Println(incrementer.inc)
}
