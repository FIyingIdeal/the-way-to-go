package main

import "fmt"

func main() {
	a := 10
	var aP *int = &a
	var b int = *aP
	fmt.Printf("ap is %p, b is %d", aP, b)
}
