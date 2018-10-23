package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			fmt.Println(incrementer, "is initial val incrementor")
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer, "is incrementor from var v")
			wg.Done()
		}()  //anon func
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}