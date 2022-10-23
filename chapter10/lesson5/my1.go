package main

import "fmt"

type Student struct {
	name string
	// int 是一个匿名字段，类型就是字段名
	int
}

func main() {
	s := Student{"zhangsan", 22}
	// 类型就是字段名，因此可以通过 s.int 访问
	age := s.int
	fmt.Println("age is :", age)
}
