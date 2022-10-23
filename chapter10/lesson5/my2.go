package main

import "fmt"

type Inner struct {
	a int
	b int
}

type Outter struct {
	c int
	Inner
}

func main() {
	outter := Outter{3, Inner{1, 2}}
	a := outter.a
	b := outter.b
	c := outter.c
	fmt.Printf("a is %d, b is %d, c is %d", a, b, c)
}
