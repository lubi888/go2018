package main

import (
	"math/rand"
	"net/http"
)

func deck(w http.ResponseWriter, r *http.Request) {
	x := 127137 //decimal of unicode chars deck cards
	a := []int{}
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
	uu := 24
	y := 37
	z := 50
	//skip 2 blank cards at time
	a = append(a[:s], a[s+2:]...)
	a = append(a[:t], a[t+2:]...)
	a = append(a[:u], a[u+2:]...)
	//remove 4 'c' noble cards
	a = append(a[:v], a[v+1:]...)
	a = append(a[:uu], a[uu+1:]...)
	a = append(a[:y], a[y+1:]...)
	a = append(a[:z], a[z+1:]...)
	//shuffle deck decimals
	rand.Shuffle(52, func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	//convert ints to strings
	d := [52]string{}
	for t := 0; t < 52; t++ {
		d[t] = string(a[t])
	}
	//fmt.Print(d, "\n\n")
	//scrolling timer
	//for t := 0; t < 52; t++ {
	//	timer:= time.NewTimer(time.Millisecond * 750)
	//	<- timer.C
	//	fmt.Println("\t\t", d[t])
	//}
	tpl.ExecuteTemplate(w, "index.gohtml", d)
}
