package main

import (
	"fmt"
	//"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	var mx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			//fmt.Println(incrementer, "is initial val incrementor")  //remove this to remove race
			mx.Lock()
			v := incrementer
			//runtime.Gosched()  //remove runtime.schedule
			v++
			incrementer = v
			fmt.Println(incrementer, "is incrementor from var v")
			mx.Unlock()
			wg.Done()
		}()  //anon func
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}

