package main

import (
	"fmt"
	"strconv"
)

type Integer int

func (i *Integer) String() string {
	return strconv.Itoa(int(*i))
}

func (i Integer) get() int {
	return int(i)
}

func main() {
	i := Integer(10)
	fmt.Println(i.String())
	fmt.Println(i.get())
}
