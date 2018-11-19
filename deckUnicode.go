package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("try some arrays with unicode")
	//127136 is card back	183 a joker
	x := 127137 //ascii    decimal
	a := []int{}
	fmt.Println(len(a)) //0
	for i := x; i <= 127198; i++ {
		a = append(a, i)
	}
	//need remove  151, 152, 167, 168, 183, 184,   14,15, 30, 31, 46,47
	fmt.Println(len(a)) //original set unicode
	//remove blank cards and jokers
	s := 14
	t := 28
	u := 42
	//remove "C" card between J-Q
	v := 11
	w := 24
	y := 37
	z := 50
	a = append(a[:s], a[s+2:]...)
	a = append(a[:t], a[t+2:]...)
	a = append(a[:u], a[u+2:]...)
	a = append(a[:v], a[v+1:]...)
	a = append(a[:w], a[w+1:]...)
	a = append(a[:y], a[y+1:]...)
	a = append(a[:z], a[z+1:]...)

	fmt.Println(len(a)) //length deck 52
	//for z := 0; z < len(a); z++ {
	//	fmt.Println(a[z])         //print a decimal   127194
	//	fmt.Printf("%U\n", a[z])  //print unicode num  U+1F0DE
	//	fmt.Printf("%#U\n", a[z]) //print unicode num + character    U+1F0DE '🃞'
	//}
	//
	//fmt.Printf("%T\n", a[0])
	//h := a[0]
	//fmt.Println(h, "is h")
	//j := fmt.Sprintf("%U", h)
	//fmt.Println(j, "is j") //U+1F0A1
	//fmt.Printf("%T\n", j)
	//fmt.Println("\\U000"+j[2:], "and some") // Print \U000f554A
	//k := fmt.Sprint("\"\\U000" + j[2:] + "\"")
	//fmt.Println(k + " is also k") //"\U0001F0A1"
	//fmt.Println(k)
	//fmt.Println("\U0001F0A1 is just a number ")
	//p := fmt.Sprintf("%#U", h)
	//fmt.Println(p) //U+1F0A1 '🂡'
	//p1 := fmt.Sprint(p[9:])
	//fmt.Println(p1)
	//
	fmt.Println(a, "is a")
	c := []string{}
	fmt.Println(c, "empty c")
	fmt.Println(len(c))
	for z := 0; z <= 51; z++ {
		t5 := fmt.Sprintf("%#U", a[z])
		t6 := fmt.Sprint(t5[9:])
		c = append(c, t6)
	}
	//fmt.Println(c)
	//shuffle a slice
	rand.Shuffle(52, func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println(a)
	fmt.Printf("%T\n", a[3])
	e := string(a[4])
	fmt.Println(e)
	d := [52]string{}
	for t := 0; t < 52; t++ {
		d[t] = string(a[t])
	}
	fmt.Print(d, "\n\n")
	for y = 0; y < 52; y++ {
		fmt.Print(d[y], "\t")
	}
	//not working y not index
	//for _, y := range d {
	//	fmt.Print(d[y], "\t")

	//rand.Shuffle(52, func(i, j int) {
	//	c[i], c[j] = c[j], c[i]
	//})

	//fmt.Println(c, "\t\t")
	//for z:=0; z < len(c); z++ {
	//	fmt.Println(c[z])
	//}
}
