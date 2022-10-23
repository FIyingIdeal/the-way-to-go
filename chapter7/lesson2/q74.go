package main

import (
	"fmt"
)

func main() {
	s1 := []byte{'p', 'o', 'e', 'm'}
	fmt.Println("s1 is ", s1)

	s2 := s1[2:]
	fmt.Println("s2 is ", s2)

	s2[1] = 't'
	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)

}
