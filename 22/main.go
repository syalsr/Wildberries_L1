package main

import (
	"fmt"
	"math/big"
)

/*
TODO Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20
*/

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}
func div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}
func multiple(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}
func subtract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}
func main() {
	a := big.Int{}
	b := big.Int{}

	a.SetInt64(100000000000000000)
	b.SetInt64(100000000000000000)

	fmt.Println(multiple(&a, &b))
	fmt.Println(div(&a, &b))
	fmt.Println(subtract(&a, &b))
	fmt.Println(add(&a, &b))
}
