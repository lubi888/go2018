package main

import (
	"fmt"
	//"math"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7}
	xi1 := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println(foo1(xi...), "foo1")
	fmt.Println(bar1(xi1), "bar1")
	defer fmt.Println(foo1(xi...), "foo1 deferred")
	fmt.Println(bar1(xi1), "bar1")

	p1 := personp{
		first: string("tom"),
		last:  string("tom"),
		age:   int(22),
	}
	p1.speak()
}
func foo() int {
	return 42
}
func bar() string {
	return "stringgy"
}
func foo1(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func bar1(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

type personp struct {
	first string
	last  string
	age   int
}

func (p personp) speak() {
	fmt.Println("hi my name is", p.first, p.last, "and I am", p.age, "years young")
	//fmt.Println(x)
}
