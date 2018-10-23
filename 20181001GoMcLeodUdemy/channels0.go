package main

import "fmt"

func main() {
	c := make(chan int)
	go func() { //wrap in goroutine to make 'pass the baton'
		c <- 42
	}()
	fmt.Println(<-c, "is channel dechannelled") //dechannel?
	//chanel buffer
	c1 := make(chan int, 1)
	c1 <- 42
	fmt.Println(<-c1, "channel with buffer")

	c2 := make(chan int, 2)
	c2 <- 42
	c2 <- 43
	fmt.Println(<-c2, "c2 multi")
	fmt.Println(<-c2, "c2 other multis")
	fmt.Println("-----")
	fmt.Printf("%T", c2)
}


