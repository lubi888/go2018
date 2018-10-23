package main

import (
	"fmt"
)

func main() {
	a := 42
	beta := "string song"
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b, "value of b")
	fmt.Println(*b, "value of *b") // * gives you the value stored at an address when you have the address
	fmt.Printf("%v", *b)           // * gives you the value stored at an address when you have the address
	fmt.Println(*&a)
	*b = 43
	fmt.Println(a)
	fmt.Println(beta)
	fmt.Println(&beta)
	fmt.Printf("%T\n", beta)
	fmt.Printf("%T\n", &beta)
	c := false
	fmt.Println(c, "is c val")
	//fmt.Println(*c)"
	fmt.Println(*&c, "is c pointer address val")
	fmt.Printf("%T\n", c)
	//fmt.Println("%T\n", *c)
	fmt.Printf("%T\n", *&c)
	fmt.Println("-------")
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foofl(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foofl(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
