package main

import "fmt"

type personP struct {
	name string
}

func changeMe(p *personP) *personP {
	(*p).name = "nevill"
	return p
}

func main() {
	v := 4
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	fmt.Println(&v)
	//fmt.Println(*v)
	//fmt.Println(&*v)
	fmt.Println(*&v)
	//go poerne in a gyre?
	//for i:=0; i<100; i++ {
	//	fmt.Println(i, //system pyrnes)
	//}
	pp := personP{
		name: string("toby mcguire"),
	}
	fmt.Println(pp)
	fmt.Println(pp.name)
	fmt.Println(&pp)
	changeMe(&pp)
	fmt.Println(pp)
	fmt.Println(pp.name)
	fmt.Println()
}
