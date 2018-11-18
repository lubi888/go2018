package main

import (
	"fmt"
	"unicode"
)

const (
	//ro	= '\\u13084'
	MaxRune = '\U0010FFFF'
)

func main() {
	const ucG = 'G'
	const T = "\u13093"
	fmt.Printf("%#U\n", unicode.ToLower(ucG))
	fmt.Printf("thry to pring anything\n")
	fmt.Println("\u007A") // single '   = ./unico.go:13: cannot use 'z' (type rune) as type string in argument to fmt.Printf
	fmt.Printf("\u1308\n")
	fmt.Printf("\u13093\n")
	//fmt.Printf('\u1309')
	//fmt.Printf("\u13084\n")
	fmt.Printf("%#U\n", unicode.Egyptian_Hieroglyphs)
	fmt.Printf("%#U\n", unicode.Coptic)
	fmt.Printf("%#U\n", '\U00013000')
	fmt.Printf("\U00013009\n")
	fmt.Printf("\U00013010\n")
	fmt.Printf("\U00013011\n")
	fmt.Printf("\U00013012\n")
	fmt.Println("\U00013015")
	fmt.Println("\U0001342E")
	fmt.Printf("\360\223\200\200")
	//fmt.Printf("\U0xF0 0x93 0x80 0x80")
	//fmt.Printf("%#U\n", unicode.U+13000)
	fmt.Printf("\360\223\202\272")
	fmt.Println("\U000130B9")
	fmt.Println("\360\223\217\207") //U+133C7  C octal escaped UTF-8
	fmt.Println("\U000133C7")       //U+133C7
	fmt.Printf("%#U\n", '\130') //prints U+0058 'X'
	//fmt.Printf("%#U\n", "\130")
	//fmt.Printf("%U\n", "\u13096")
	//	fmt.Printf({[U+13000]}, U+13000})
	//fmt.Printf("%#U\n", "U+0058")
	//fmt.Printf("%#U\n", "U+13000") //	fmt.Printf("%#U\n", "U+13006")
	//fmt.Printf("%#U\n", (T))
	fmt.Printf("%#U\n", unicode.Ogham)
	fmt.Printf("\u1685\n")
	fmt.Printf("\u1696\n")
	//fmt.Printf("\unicode.Egyptian_Hieroglyphs.R16.13093\n")
	//fmt.Printf("\unicode.Egyptian_Hieroglyphs.13093")
	//fmt.Printf("\unicode.Egyptian_Hieroglyphs.13090\n")

	//2018.10.05
	//77824	decimal & html		0x13000	U+13000 '𓀀'
	fmt.Println("try some arrays with unicode")
	x := 77824 //ascii a
	a := []int{}
	//a := make([]int,1)
	fmt.Println(len(a)) //0
	for i := x; i <= 77824+1070; i++ { //1072 hiero set
		a = append(a, i)
	}
	fmt.Println(len(a))

	for z := 0; z < len(a); z++ {
		fmt.Print(a[z], "\tdecimal & html")
		fmt.Printf("\t\t%#x\t%#U\n", a[z], a[z])
		//fmt.Printf("%#U\t", a[z])
		//fmt.Printf("%#U\n", a[z])
	}
	for z := 0; z < len(a); z++ {
		fmt.Printf("%#U", a[z])
	}
}

