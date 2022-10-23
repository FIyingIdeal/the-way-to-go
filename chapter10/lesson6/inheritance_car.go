package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

// 当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，类似于Java中的继承
type Car struct {
	wheelCount int
	Engine
}

func (c Car) numberOfWheels() int {
	return c.wheelCount
}

func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Merkel")
}

func main() {
	m := Mercedes{Car{4, nil}}
	fmt.Println("A meredes has this many wheels : ", m.wheelCount)

	m.GoToWorkIn()
	m.sayHiToMerkel()
}
