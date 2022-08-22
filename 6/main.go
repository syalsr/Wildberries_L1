package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
TODO Реализовать все возможные способы остановки выполнения горутины.
*/

func withcancel() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	go gogo(ctx, wg)
	time.Sleep(2 * time.Second)
	cancel() //по истечению 2 секунд горутина закончит свою работу

	wg.Wait()
}
func withtimeout() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go gogo(ctx, wg) //по истечению 5 секунд горутина закончит свою работу

	wg.Wait()
}

func gogo(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("working")
		}
	}
}

func canceL() {
	quit := make(chan bool)
	go func() {
		for {
			_, ok := <-quit
			if !ok { //если канал закрыт выходим из горутины
				fmt.Println("!Ok")
				return
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(quit)
	time.Sleep(2 * time.Second)
}

func main() {

}
