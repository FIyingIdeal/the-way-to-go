package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	// s1 为 student 的变量，此时也会为 s1 分配内存，并零值化内存
	var s1 student
	fmt.Println("Before set s1 value : ", s1)
	s1.name = "zhangsan"
	s1.age = 10
	fmt.Println("Ater set s1 value : ", s1)

	// 更简短的变量声明方式
	var s11 student = student{"lisi", 20}
	fmt.Println("Ater set s11 value : ", s11)

	// 更简短的变量声明方式
	s12 := student{age: 26, name: "lili"}
	fmt.Println("After set s12 value : ", s12)

	// s2 是指向一个结构体类型变量的指针，
	var s2 *student = new(student)
	s2.name = "wangwu"
	s2.age = 25
	fmt.Println("Before set s2 value : ", s2)

	// 更简短的指针声明方式，只是 new() 的一种简写，s22 的类型是 *student
	s22 := &student{"zhaoliu", 30}
	fmt.Println("Before set s22 value : ", s22)
}
