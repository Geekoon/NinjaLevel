package main

import (
	"fmt"
)

const (
	utc = 67
	tc int = 23
)

var a int

func main() {
	a = 35
	fmt.Printf("%b\n", a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%#x\n", a)
	fmt.Println(utc)
	fmt.Println(tc)
	b := a << 1
	fmt.Printf("%b\n", b)
	fmt.Printf("%d\n", b)
	fmt.Printf("%#x\n", b)
}
