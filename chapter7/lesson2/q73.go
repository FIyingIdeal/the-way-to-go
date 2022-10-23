package main

import "fmt"

func main() {
	s := make([]int, 5)
	// 5
	fmt.Println("len(s) is ", len(s))
	// 5
	fmt.Println("cap(s) is ", cap(s))

	s = s[2:4]
	// 2
	fmt.Println("len(s) is ", len(s))
	// 3
	fmt.Println("cap(s) is ", cap(s))

}
