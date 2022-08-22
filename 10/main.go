package main

import "fmt"

/*
TODO Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func round(n float64) int64 {
	return int64(n/10) * 10
}

func groupTemperature(arr []float64) map[int64][]float64 {
	m := make(map[int64][]float64)
	for _, item := range arr {
		i := round(item) //округляем до десятых
		m[i] = append(m[i], item)
	}
	return m
}

func main() {
	m := groupTemperature(
		[]float64{
			1.0, 2.0, 3.5, 10.9, 11.3, 13.0, 43.6, 41.0, 49.0,
		},
	)
	fmt.Println(m)
}
