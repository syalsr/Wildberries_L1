package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

type Map struct {
	ma    map[int]int
	mutex *sync.Mutex
}

func (m *Map) AddElem(key, value int) {
	m.mutex.Lock()         //горутина захватывает мютекс
	defer m.mutex.Unlock() //при выходе из функции - освобождает
	m.ma[key] = value
}

func CreateMap() *Map {
	return &Map{ma: make(map[int]int), mutex: new(sync.Mutex)}
}

func main() {
	Map := CreateMap()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 10; i < 20; i++ {
			Map.AddElem(i, -i)
		}

	}()
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			Map.AddElem(i, -i)
		}
	}()

	wg.Wait()

	for key, value := range Map.ma {
		fmt.Printf("Key: %d - Value: %d\n", key, value)
	}
}
