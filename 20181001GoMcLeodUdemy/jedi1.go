package main

import (
	"fmt"
	"math/rand"
)

const (
	yer   = iota
	yer19 = 2018 + iota
	yer20 = 2018 + iota
	yer21 = 2018 + iota
)

//const nums []string(
//ace = iota
//two
//three
//)
//
//const (
//hearts = iota
//diamonds
//clubs
//spades
//)

func main() {
	//x := 42
	//y := `that's the answer`
	//z := true
	//
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)
	//
	//fmt.Println(x, y, z)
	//fmt.Printf("%T", "%T", "%T", x, y, z)
	//fmt.Println()
	//fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	//fmt.Println(yer19, yer20, yer21)

	face := [13]string{
		"ace",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"jack",
		"queen",
		"king",
	}

	suit := [4]string{
		//"hearts",
		//"diamonds",
		//"clubs",
		//"spades",

		"\u2660", //hearts
		"\u2663", //diamonds
		"\u2665", //clubs
		"\u2666", //spades
	}

	//"spades":   "\u2660",
	//"clubs":    "\u2663",
	//"hearts":   "\u2665",
	//"diams":    "\u2666",
	//
	//suit := map[string]string{
	//	"spades": "\u2660",
	//	"clubs":  "\u2663",
	//	"hearts": "\u2665",
	//	"diams":  "\u2666",
	//}

	fmt.Println("num of suit")
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(face[i])
			fmt.Print("\tof\t")
			fmt.Print(suit[j])
			fmt.Println("")
		}
		fmt.Println("")
	}

	fmt.Println("-------------------suit num")
	for j := 0; j < 4; j++ {
		for i := 0; i < 13; i++ {
			fmt.Print(suit[j], " ")
			fmt.Print(face[i], "\n")
		}

	}

	rn := rand.Intn(13)
	rn1 := rand.Intn(13) + 1
	fmt.Println(rn + 1)
	fmt.Println(rn1)

	rs := rand.Intn(3)
	fmt.Println(rs + 1)

	fmt.Println("num of suit")
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(face[i])
			fmt.Print("\tof\t")
			fmt.Print(suit[j])
			fmt.Println("")
		}
		fmt.Println("")
	}

	fmt.Println("how about typing stdio? a new card: ")

	var i int
	fmt.Scan(&i)

	for i == 0 {
		fmt.Println("that worked")
		break
		//fmt.Scan(&i)
	}

	fmt.Println("-------------------suit num")
	//for j := 0; j < 4; j++ {
	//	for i := 0; i < 13; i++ {
	fmt.Print(suit[rs], " ")
	fmt.Print(face[rn], "\n")
	//	}
	//
	//}
}
