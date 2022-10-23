package main

import "fmt"

func main() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println("b[1:4] is ", b[1:4])
	fmt.Println("b[:2] is ", b[:2])
	fmt.Println("b[2:] is ", b[2:])
	fmt.Println("b[:]  is ", b[:])
}
