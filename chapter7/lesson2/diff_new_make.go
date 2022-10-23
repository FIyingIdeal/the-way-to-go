package main

import "fmt"

// 理解 new 和 make 的区别

func main() {
	// 注意：这里是 new 一个数组，返回数组的指针，不是一个切片
	var p *[]int = new([]int)
	// p is  &[]
	fmt.Println("p is ", p)
	// *p is nil? true    表明数组没有初始化
	fmt.Println("*p is nil?", *p == nil)

	// 注意：这里返回一个 slice，slice本身就是一个引用类型
	var s []int = make([]int, 0)
	// s is nil? false    表明数组已经初始化，只不过是个长度为0的空数组
	fmt.Println("s is nil?", s == nil)
}
