package main

import (
	"fmt"
)

type person struct {
	first string
	second string
	age int
}

func (p *person) speak() {
	fmt.Println(p.first, "is speaking")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person {
		first: "James",
		second: "Bond",
		age: 35,
	}
	fmt.Println("hi", p1)
	saySomething(&p1)
}