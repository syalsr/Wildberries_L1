package main

import (
	"time"
)

/*
TODO Реализовать собственную функцию sleep
*/

func sleep1(duration time.Duration) {
	<-time.After(duration) //спустя duration секунд продолжит работу, выйдет из функции
}
func slee2(duration time.Duration) {
	<-time.NewTimer(duration).C
}
func main() {
	sleep1(2 * time.Second)
}
