package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	// 结构体可以包含一个或多个匿名（或内嵌）字段
	// 此时类型就是字段的名字，但也因此一个结构体中对于每一种数据类型只能有一个匿名字段
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 1
	outer.c = 4.2
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	outer2 := outerS{6, 4.2, 60, innerS{5, 10}}
	fmt.Println("outer2 is ", outer2)
}
