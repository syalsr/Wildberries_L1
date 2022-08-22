package main

import "fmt"

/*
TODO Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human
(аналог наследования).
*/

type Human struct {
	name string
	age  uint
}

func (human *Human) sayHello() {
	fmt.Printf("Hello, I am %s\n", human.name)
}
func (human *Human) sayAge() {
	fmt.Printf("Hello, I am %d\n", human.age)
}

type Action struct {
	Human
	job string
}

/*
Встроили анонимное поле, теперь можем обращаться к данным и методам структуры human через
action, в случае если мы переопределяем данные или методы в структуре Action, и хотим обратиться
к Human нужно делать уже так action.Human.age = 5
*/

func (action *Action) sayHello() {
	fmt.Printf("Hello, I am %s\n", action.name)
}
func (action *Action) sayAge() {
	fmt.Printf("Hello, I am %d\n", action.age)
}
func (action *Action) sayJob() {
	fmt.Printf("Hello, I am %d\n", action.job)
}
func main() {
	h := Human{name: "Alex", age: 32}
	h.sayAge()
	h.sayHello()

	a := Action{job: "Freelance", Human: Human{name: "Xela", age: 12}}
	a.sayHello()
	a.sayAge()
}
