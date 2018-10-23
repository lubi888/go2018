package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	//buffered
	buff := make(chan int, 1)

	buff <- 42

	fmt.Println(<-buff, "is buffered chan")
}
