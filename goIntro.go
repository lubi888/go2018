package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x, y float64) float64 { // ( x int, y int) => (x, y, int)
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

//var c, python, java bool
var i, j int = 1, 3

func main() {
	fmt.Println("welcome to the playground of go and dart.")
	fmt.Println("the time is ", time.Now())
	fmt.Println("my fav num is", rand.Intn(100))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	fmt.Println("there are other %g issues. this is println", math.Sqrt(9)) //, math.SqrtPi)
	fmt.Printf("what about other %g notes?", math.Sqrt(6))
	fmt.Println()
	fmt.Println("give me ln v. f, ", cmplx.Tan(90))
	fmt.Print("another complex, ", cmplx.Sqrt(9))
	fmt.Println()
	fmt.Print("this is pi, ", math.Pi, " ,->")
	fmt.Printf("this is pi, %g", math.Pi) // without %g == %!(EXTRA float64=3.141592653589793)    & no new line
	fmt.Println("this is pi, ", math.Pi, "shove what?")
	fmt.Println("this is phi, ", math.Phi)
	fmt.Print(" this is also phi with a print, not println, and without a modulusg ", math.Phi)
	fmt.Println()
	fmt.Println(" this is also phi with a println, not printf, and without a modulusg ", math.Phi)
	fmt.Printf(" this is also phi with a printf, not println, and without a modulusg ", math.Phi) //%!(EXTRA float64=1.618033988749895)
	fmt.Println()
	fmt.Printf(" and antoher phi with modg %g", math.Phi) //"thus avoiding %!(EXTRA float64=1.618033988749895)")
	fmt.Println()
	fmt.Printf(" and antoher phi with modg %g", math.Phi, "thus what? == %!(EXTRA string=xx)")
	fmt.Println()

	fmt.Printf(" and antoher phi with modg %g", math.Phi, "thus avoiding (EXTRA float64=1.618033988749895)")
	fmt.Println()
	fmt.Printf(" and antoher phi with modg %g", math.Phi, "thus avoiding %!(EXTRA float64=1.618033988749895)")
	fmt.Println(add(25, 25))
	fmt.Printf("25 added to itself is %g", add(25, 25))
	//fmt.Println()
	fmt.Println(" ", add(25, 25)/2, " ")
	fmt.Println(add(25, 25) * 25)
	fmt.Println(add(25, math.Sqrt(9)))

	a, b := swap("hello", "go")
	fmt.Println(a, b)
	fmt.Println(b, a, a, b)
	//	fmt.Print("the split of 18 is: ", split(18))
	fmt.Print("the split of 9 is ")
	fmt.Println(split(9))
	//	var i int
	var c, python, java = true, false, "no falsch!"
	k := 1
	fmt.Println(i, j, c, python)
	fmt.Println(java, k)

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
		s      string     = "what"
	)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe) //bool(false)
	fmt.Println('f', f, ToBe, ToBe)
	fmt.Printf(f, MaxInt) //, MaxInt) // uint64(18446744073709551615)  uint64(%!v(MISSING))
	//	fmt.Printf(z, z, f) //(z, f, f)  == / cannot use z (type complex128) as type string in argument to fmt.Printf
	fmt.Printf(f, z, z) // = complex128((2+3i))
	//fmt.Printf(f, f, f) //= string(%T(%v)
	//fmt.Println(f, z, z) // %T(%v) (2+3i) (2+3i)
	fmt.Printf(f, s, s)
}

