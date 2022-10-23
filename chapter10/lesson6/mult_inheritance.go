package main

import "fmt"

// 通过多个内嵌类型实现多继承

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type CameraPhone struct {
	Phone
	Camera
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors ...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}
