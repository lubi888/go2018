package main

import "fmt"

type personl struct {
	first string
	last  string
	l     []string
}

func main() {

	p1 := personl{
		first: string("todd"),
		last:  string("mccleod"),
		l:     []string{"nlk", "tty", "wwe", "ssw", "llo"},
	}
	p12 := personl{
		first: string("dopt"),
		last:  string("mc"),
		l:     []string{"nlk", "tty", "wwe", "ssw", "llo"},
	}
	fmt.Println(p1)
	fmt.Println(p1.last, p1.l[3])
	//fmt.Println(
	for i, v := range p1.l {
		fmt.Println(i, v)
	}
	//)
	fmt.Println(p12)

	//make a map from a struct
	m := map[string][]string{
		//lastname: first
		p1.last: {p1.l[1], p1.l[2]},
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
		for k1, v1 := range p1.l {
			fmt.Println(k1, v1)
		}
	}

}
