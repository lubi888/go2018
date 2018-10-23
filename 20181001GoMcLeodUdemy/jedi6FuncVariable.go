package main

import "fmt"

func double(x float64) float64 {
	return x * x
}

func main() {
	t := double(4.3)
	fmt.Println(t, "is the func var")
	g := double(6)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T,%v", double(5), double(5))
}
