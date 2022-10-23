package main

import "fmt"

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(id int) {
	b.id = id
}

type Person struct {
	Base
	FirstName string
	LaseName  string
}

type Employee struct {
	Person
	salary int
}

func main() {
	employee := &Employee{Person{Base{1}, "zhang", "san"}, 100}
	fmt.Println(employee.id)

	employee.SetId(3)
	fmt.Println(employee.id)
}
