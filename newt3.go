package main

import (
	"fmt"
	"math"
)

func loop(x float64) string {
	for x < 25 {
		x += 1 //x = x + 1
		fmt.Print("the root of ")
		fmt.Print(x)
		fmt.Print(" is ")
		fmt.Println(math.Sqrt(x))
	}
	return fmt.Sprintln("--loop roots 25 to string--")
}

func newt(r float64, g float64) (ans float64) {
	fmt.Println("guess for", r, "was =", g)
	//neg root answer not using Newt's Method
	//make neg root function availbe to all funcs(newt, newtCube, newtQuad, newtQuinque)?
	if r <= 0 {
		q := math.Sqrt(-r)
		fmt.Printf("%g has an imaginary root of %gi", r, q)
		fmt.Println()
		return q
	}
	//check lucky guess
	if z := math.Pow(g, 2); z == r {
		fmt.Printf("%g has a positive root of %g", r, g)
		fmt.Println()
		return g
	}
	//newton's interations
	for i := 0; i < 10; i++ {
		z := g
		delta := .00000001
		g = g - ((g*g - r) / (2 * g))
		fmt.Println("root guess for", r, "is now", g, "at", i+1, "newton iterations")

		if math.Abs(z-g) < delta {
			fmt.Println("delta is now < .00000001")
			return g // this return stops iterations i
		}
	}
	return //naked return because returned float64 is named variable? otherwise 'return g'
}

func newtCube(r float64, g float64) (anscube float64) {
	fmt.Println("cube root guess for", r, "was =", g)
	for i := 0; i < 10; i++ {
		z := g
		delta := .00000001
		//g = g - ((g*g*g - r) / (3 * g * g))      for cube root
		g = g - ((math.Pow(g, 3) - r) / (3 * math.Pow(g, 2)))
		fmt.Println("cube root guess for", r, "is now", g, "at", i+1, "newton iterations")

		if math.Abs(z-g) < delta {
			fmt.Println("delta is now < .00000001")
			return g // this return stops iterations 10
		}
	}
	return
}

func newtQuad(r float64, g float64) (ansquad float64) {
	fmt.Println("cube root guess for", r, "was =", g)
	for i := 0; i < 10; i++ {
		z := g
		delta := .00000001
		//g = g - ((g*g*g*g - r) / (4 * g * g * g))     for 4th root
		g = g - ((math.Pow(g, 4) - r) / (4 * math.Pow(g, 3)))
		fmt.Println("4th root guess for", r, "is now", g, "at", i+1, "newton iterations")

		if math.Abs(z-g) < delta {
			fmt.Println("delta is now < .00000001")
			return g // this return stops iterations 10
		}
	}
	return
}

func newtQuinque(r float64, g float64) (ansquinque float64) {
	fmt.Println("5th root guess for", r, "was =", g)

	for i := 0; i < 20; i++ {
		z := g
		delta := .00000001
		//g = g - ((g*g*g*g*g - r) / (5 * g * g * g * g))
		g = g - ((math.Pow(g, 5) - r) / (5 * math.Pow(g, 4)))
		fmt.Println("5th root guess for", r, "is now", g, "at", i+1, "newton iterations")

		if math.Abs(z-g) < delta {
			fmt.Println("delta is now < .00000001")
			return g // this return stops iterations 10
		}
	}
	return
}

func main() {
	fmt.Print(math.Sqrt(81))
	fmt.Println(" is the root of 81")
	fmt.Println("--using math.Sqrt library method--")
	fmt.Print(loop(0))
	fmt.Println(newt(1, 1))
	fmt.Println(newt(9, 3))
	fmt.Print(newt(-9, -3)) //neg roots
	fmt.Println("i")
	fmt.Println("---end of neg imaginary root using math.Sqrt(-r) not Newtons Method---")
	fmt.Println("--end of if statements before loops--")
	fmt.Println("now need to input 2 numbers; the radicand and a guess that is the radicand's square root.  //for future io")
	fmt.Println("could also make a simple guess based on radicand/2 or another simp guess? this would reduce need for 2 inputs")
	fmt.Println("--api?--")
	fmt.Println("---------")
	fmt.Println("--try number--")
	fmt.Println(newt(2, 1))
	fmt.Println("--square root by num float64--")
	fmt.Print(newt(25, 3)) // radi2cand and guess
	fmt.Println(" is the answer")
	fmt.Println("--root25--")
	fmt.Println("---try higher powers roots---")
	fmt.Println(newtCube(27, 2))
	fmt.Println(newtQuad(16, 1.5))
	fmt.Println(newtQuinque(20000000, 20))
}

//https://tour.golang.org/flowcontrol/8
//https://en.wikipedia.org/w/index.php?title=Newton%27s_method&oldid=731108767

//https://mitpress.mit.edu/sicp/chapter1/node9.html
//https://mitpress.mit.edu/sicp/chapter1/node9.html#exexsqrtendtest
