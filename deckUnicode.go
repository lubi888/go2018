package main

import "fmt"

func main() {
	//127136 is card back	183 a joker
	x := 127137 //ascii    decimal
	a := []int{}
	colour := bool(false)
	for i := x; i <= 127198; i++ {
		a = append(a, i)
	}
	//need remove  151, 152, 167, 168, 183, 184,   14,15, 30, 31, 46,47
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

	//fmt.Println(len(a)) //length deck 52
	//for z := 0; z < len(a); z++ {
	//	fmt.Println(a[z])         //print a decimal   127194
	//	fmt.Printf("%U\n", a[z])  //print unicode num  U+1F0DE
	//	fmt.Printf("%#U\n", a[z]) //print unicode num + character    U+1F0DE '🃞'
	//}

	//Give colour to cards
	//struct { []a{int}, color true bool}
	//127137 Ace Spades   -50
	//127153 ace hearts		-66
	//127169 ace diamonds	-82
	//127185 ace trumps		-98

	//Non-integer slice index z less... (Ctrl+F1)
	//Inspection info: Reports invalid index and slice expression
	//for z := [52]; z < 52; z++ {
	//	if a[z] >127150 || a[z] <127185 {
	//		colour = true
	//	} else {
	//		colour = false
	//	}
	//}

	//shuffle a slice
	//rand.Shuffle(52, func(i, j int) {
	//	a[i], a[j] = a[j], a[i]
	//})
	////create new slice deck strings
	//d := [52]string{}
	//for t := 0; t < 52; t++ {
	//	d[t] = string(a[t])
	//}

	//print random deck slice
	//fmt.Print(d, "\n\n")
	//for y = 0; y < 52; y++ {
	//	fmt.Print(d[y], "\t")
	//}

	//scrolling timer
	for t := 0; t < 52; t++ {
		//timer:= time.NewTimer(time.Millisecond * 1000)
		//<- timer.C
		//diplay vert or horiz
		//fmt.Print(a[t])
		fmt.Print("\t", a[t])   //z[t]
			//if a[z] >127150 && a[z] <127185 {
			if a[t] >127137 && a[t] <127151 {
				colour = true
				fmt.Print(a[z])
				fmt.Println(colour)
			}	else {
				colour = false
				fmt.Print(a[z], colour, "\t")
			}
		//fmt.Print("\t", colour)
		//fmt.Println()
	}
}
