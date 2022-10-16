package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	printValues()

	numx2, numx3 = getX2AndX3_2(num)
	printValues()
}

func getX2AndX3(num int) (int, int) {
	return 2 * num, 3 * num
}

func getX2AndX3_2(num int) (x2, x3 int) {
	x2 = 2 * num
	x3 = 3 * num
	return
}

func printValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}
