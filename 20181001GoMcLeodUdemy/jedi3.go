package main

import "fmt"

func main() {
	fmt.Println("hi")
	s := 65
	t := 90
	a := []int{}
	fmt.Println(len(a))
	for z := s; z <= t; z++ {
		a = append(a, z)
	}
	fmt.Println(len(a))

	for g := 0; g < len(a); g++ {
		fmt.Printf("%v\n", a[g])
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", a[g])
		}
	}
	fmt.Println("blah")
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Print(i, "\t")
		}
	}
}
