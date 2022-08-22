package main

import "fmt"

/*
TODO Дана переменная int64. Разработать программу которая устанавливает
i-й бит в 1 или 0
*/

/*
b: false - bit 0; true - bit 1
pos: 0 - 63
*/
func setBit(num int64, pos uint8, b bool) (n int64) {
	//с помощью сдвига устанавилваем в pos bit 1
	if b {
		n = num | 1<<pos //побитовое ИЛИ - если значение = 1, то будет 1; если 0, то будет 1
	} else {
		n = num &^ (1 << pos) //побитовое И НЕ
	}
	return
}

func main() {
	fmt.Println(setBit(255, 7, false))
}
