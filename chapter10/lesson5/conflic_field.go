package main

import "fmt"

type A struct{ a int }
type B struct{ a, b int }

type C struct {
	A
	B
	b float32
}

func main() {
	c := C{A{11}, B{21, 22}, 3.1}
	// 报错：ambiguous selector c.acompiler
	// 不知道是 c.A.a 还是 c.B.a
	//aa := c.a

	// 可以通过加匿名字段名的方式明确指定要用哪个结构体内的字段
	aa := c.A.a
	// 11
	fmt.Println("aa is ", aa)

	ba := c.B.a
	// 21
	fmt.Println("ba is ", ba)

	// 外层结构体C中的字段b覆盖了内层匿名结构体B中的字段b
	cb := c.b
	// 3.1
	fmt.Println("bb is ", cb)

	bb := c.B.b
	// 22
	fmt.Println("bb is ", bb)
}
