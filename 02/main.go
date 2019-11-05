package main

import (
	"fmt"
)

const (
	utc = 67
	tc int = 23
)

const (
	_ = (iota + 2019)
	n1
	n2
	n3
	n4
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
	s := `dskjfn  sdf "dsfsdf" 
			sdf sdf sssdf 23r 23r df
fds f230u `
	fmt.Println(s)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}
