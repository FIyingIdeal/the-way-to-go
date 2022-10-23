package main

import (
	"bytes"
	"fmt"
)

// https://nitter.kylrth.com/wusiwan1003?cursor=HBaCwKDFpe%2FpsysAAA%3D%3D
func main() {
	buffer := bytes.NewBufferString("hello world")
	bPrefix, bSuffix := buffer.Bytes()[:4], buffer.Bytes()[4:]
	fmt.Println("bPrefix is ", string(bPrefix), ", bSuffix is ", string(bSuffix))
}
