package main

import (
	"fmt"
	//"unicode"
)

func main() {
	x := 97 //ascii a
	a := []int{}
	//a := make([]int,1)
	fmt.Println(len(a))
	for i := x; i < 123; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(len(a))

	for z := 0; z < len(a); z++ {
		fmt.Print(a[z])
		fmt.Printf("\t\t%#x\t", a[z])
		fmt.Printf("%#U\n", a[z])
	}
	//x := 97 //ascii a
	//a := []int{}
	////a := make([]int,1)
	//fmt.Println(len(a))
	//for i := x; i < 123; i++ {
	//	a = append(a, i)
	//}
	//fmt.Println(a)
	//fmt.Println(len(a))
	//
	//for z := 0; z < len(a); z++ {
	//	fmt.Print(a[z])
	//	fmt.Printf("\t\t%#x\t", a[z])
	//	fmt.Printf("%#U\n", a[z])
	//}
}
