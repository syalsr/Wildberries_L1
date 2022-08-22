package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
TODO Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала
и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(
	context context.Context, countWorkers int, chanData chan string, wg *sync.WaitGroup,
) {
	for {
		select {
		case elem := <-chanData: //берем значение из канала
			fmt.Printf(
				"Worker %d: %s\n",
				countWorkers,
				elem,
			)
		case <-context.Done(): //проверяем не завершился ли контекст
			defer wg.Done()
			fmt.Printf("CTRL+C\n")
			return
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}

	chanData := make(chan string)
	go func() {
		for {
			chanData <- "data"
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var countWorkers int
	fmt.Scan(&countWorkers)

	wg.Add(countWorkers)

	go func() {
		for countWorkers > 0 { // Запускаем N воркеров
			go worker(ctx, countWorkers, chanData, wg)
			countWorkers--
		}
	}()

	sigterm := make(chan os.Signal)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-sigterm: //при получении сигнала, отменяем контекст
			cancel()
		}
	}()

	wg.Wait()
}
