package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

type antor int

var a antor
var r int

func main() {
	fmt.Println("Hello")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	a = 42
	fmt.Printf("My Type is %T\n", a)
	fmt.Println("The meaning of my Type is", a)
	r = int(a)
	fmt.Println("The meaning of r is", r)
	fmt.Printf("%T\n", r)
}
