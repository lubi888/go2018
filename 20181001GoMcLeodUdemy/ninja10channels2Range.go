package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	//add range
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)
	//create goroutine
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}




//starting code

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	c := gen()
//	receive(c)
//
//	fmt.Println("about to exit")
//}
//
//func gen() <-chan int {
//	c := make(chan int)
//
//	for i := 0; i < 100; i++ {
//		c <- i
//	}
//
//	return c
//}
