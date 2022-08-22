package main

import (
	"context"
	"fmt"
	"time"
)

/*
TODO Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func withcancel() {
	chanData := make(chan string)
	go func() {
		for {
			chanData <- "Data"
		}
	}()

	var duration int64
	fmt.Scan(&duration)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Duration(duration) * time.Second) //по истечению времени отменяем контекст
		cancel()
	}()

	for {
		select {
		case elem := <-chanData:
			fmt.Printf("%s", elem)
		case <-ctx.Done(): //при отмене контекста возвращаемся из функции
			return
		}
	}
}

func withdeadline() {
	chanData := make(chan string)
	go func() {
		for {
			chanData <- "Data"
		}
	}()

	var duration int64
	fmt.Scan(&duration)

	ctx, _ := context.WithDeadline(
		context.Background(),
		time.Now().Add(time.Duration(duration)*time.Second),
	) //отменяем контекст через несколько секунд

	for {
		select {
		case elem := <-chanData:
			fmt.Printf("%s", elem)
		case <-ctx.Done(): //контекст отменен, возвращаемся из функции
			return
		}
	}
}
func withtimeoute() {
	chanData := make(chan string)
	go func() {
		for {
			chanData <- "Data"
		}
	}()

	var duration int64
	fmt.Scan(&duration)

	ctx, _ := context.WithTimeout(
		context.Background(),
		time.Duration(duration)*time.Second,
	) //отменяем контекст через несколько секунд

	for {
		select {
		case elem := <-chanData:
			fmt.Printf("%s", elem)
		case <-ctx.Done(): //контекст отменен, возвращаемся из функции
			return
		}
	}
}
func main() {
	withdeadline()
}
