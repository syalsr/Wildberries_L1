package main

import (
	"fmt"
	"reflect"
)

/*
TODO Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func checkVar(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	}
	return "Unknown"
}

func yetAnotherCheckVar(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func yetYetAnotherCheckVar(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func yetYetYetAnotherCheckVar(v interface{}) string {
	return reflect.ValueOf(v).Kind().String()
}

func main() {
	fmt.Println(yetYetYetAnotherCheckVar(12))
}
