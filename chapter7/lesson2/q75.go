package main

import (
	"bytes"
	"fmt"
)

func main() {
	sl := make([]byte, 4)
	fmt.Println("sl is ", sl)
	result := Append(sl, []byte{'a', 'b', 'c', 'd', 'e', 'f'})
	fmt.Println("result is ", result)
}

func Append(slice, data []byte) []byte {
	if slice == nil {
		slice = make([]byte, 4)
	}
	buffer := bytes.NewBuffer(slice)
	buffer.Write(data)
	return buffer.Bytes()
}
