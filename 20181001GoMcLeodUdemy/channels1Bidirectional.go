package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	c1 := make(chan int)
	// send
	go foo9(c1)
	// rec1eive
	bar9(c1)
	fmt.Println("about to exit")

	//e.g. of range over channels
	c4 := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c4 <- i
		}
		close(c4)
	}()

	// receive
	for v := range c4 {
		fmt.Println(v)
	}
}

// send
func foo9(c1 chan<- int) {
	c1 <- 42
}

// rec1eive
func bar9(c1 <-chan int) {
	fmt.Println(<-c1, "is bar through foo method")

}
