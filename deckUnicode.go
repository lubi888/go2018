package main

import "fmt"

func main() {
	fmt.Println("try some arrays with unicode")
	//127136 is card back	183 a joker
	x := 127137 //ascii    decimal
	a := []int{}
	fmt.Println(len(a))                //0
	for i := x; i <= 127198; i++ { 
		a = append(a, i)
	}
	//need remove  151, 152, 167, 168, 183, 184,   14,15, 30, 31, 46,47
	fmt.Println(len(a))	//original set unicode
	//remove blank cards and jokers
	s:=14
	t:=28
	u:=42 
	//remove "C" card between J-Q
	v:=11
	w:=24
	y:=37
	z:=50
	a = append(a[:s],a[s+2:]...)
	a = append(a[:t],a[t+2:]...)	
	a = append(a[:u],a[u+2:]...)	
	a = append(a[:v],a[v+1:]...)	
	a = append(a[:w],a[w+1:]...)	
	a = append(a[:y],a[y+1:]...)	
	a = append(a[:z],a[z+1:]...)	
		
	fmt.Println(len(a))  //length deck 52
	//for z := 0; z < len(a); z++ {
	//	fmt.Print(a[z], "\tdecimal & html ", z)
	//	fmt.Printf("\t\t%#x\t%#U\n", a[z], a[z])
	//}	
	for z := 0; z < len(a); z++ {
		fmt.Println(a[z])    //print a decimal
		fmt.Printf("%U\n",a[z])  //print unicode num
		fmt.Printf("%#U\n",a[z]) //print unicode num + character
	}
}
