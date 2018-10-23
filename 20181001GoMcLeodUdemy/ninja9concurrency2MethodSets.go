package main

import (
	"fmt"
	//"sync"
)

type person18 struct {
	first string
}

func (p *person18) speak() {
	//return p.first]
	fmt.Println("hello speaks")
}

type human9 interface {
	speak()
}

func saySomething7(h human9) {
	h.speak()
}

func main() {
	p1 := person18{"terry"}
	//fmt.Println(p1)
	saySomething7(&p1)
}
