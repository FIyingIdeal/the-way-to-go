package main

import "fmt"

type Log struct {
	message string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.message += "\n" + s
}

func (l *Log) String() {
	fmt.Println(l.message)
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.message = "1 - Yes we can!"

	// 简化上边的代码
	// c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}
