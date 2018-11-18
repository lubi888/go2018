package main

import "fmt"

func main() {
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
}
