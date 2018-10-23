package main

import (
	"fmt"
	//"text/scanner"
)

func main() {
	fmt.Println("how about typing something stdio?: ")

	var i int
	fmt.Scan(&i)

	for i == 0 {
		fmt.Println("try again with a digi")
		fmt.Scan(&i)
	}

	fmt.Println("read number", i, "from stdin")
	fmt.Println("how about typing something txttual stdio?: ")

	var txt string
	fmt.Scan(&txt)
	fmt.Println("what we got was ", txt, " from input")
}
