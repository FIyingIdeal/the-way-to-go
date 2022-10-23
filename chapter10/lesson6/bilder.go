package main

import (
	"fmt"
	"unsafe"
)

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main() {
	f := NewFile(3, "./test.txt")
	// 获取结构体占用了多少内存
	size := unsafe.Sizeof(f)
	fmt.Println(size)
}
