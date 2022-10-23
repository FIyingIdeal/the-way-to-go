package main

import (
	"fmt"
	"time"
)

// 如果想给 time.Time 这个 go 原生库的struct添加方法，必须自定义类型，否则会报错
// 类型和作用在它上面定义的方法必须在同一个包里定义
type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	fmt.Println("Full time now : ", m.String())

	fmt.Println("First 3 chars : ", m.first3Chars())
}
