package main

import "fmt"

func sumIt(s, t int) int {
	return s + t
}

func main() {
	x := sumIt(1, 2)
	fmt.Println(x)
}
