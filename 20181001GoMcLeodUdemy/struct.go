package main

import "fmt"

type pers struct {
	first  string
	second string
	age    int
	tel    int32
}

type sa struct {
	pers
	ltk bool
}

func main() {

	sa1 := sa{
		pers: pers{
			first:  string("jeff"),
			second: string("rey"),
			age:    int(23),
			tel:    int32(232323),
		},
		ltk: true,
	}

	p2 := pers{
		first:  string("ms"),
		second: string("moneypenny"),
		age:    int(44),
		tel:    int32(43434343),
	}
	fmt.Println(sa1)
	fmt.Println(p2.second, p2.first, p2.age)
}
