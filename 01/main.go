package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

type antor int

var a antor

func main() {
	fmt.Println("Hello")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	s = fmt.Sprintf("%v", a)
	fmt.Println(s)
}
