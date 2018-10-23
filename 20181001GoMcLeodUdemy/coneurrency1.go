package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
//func init()  {
//	//get code setup before main.
//}

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("Root\t\t", runtime.GOROOT())
	fmt.Println("Memory stats\t\t", runtime.ReadMemStats)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	if runtime.NumCPU() > 4{
		fmt.Println("we have more than 4 cpu's")
	}

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 4; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() //make sure all added and waitgroup done
}

func bar() {
	for i := 0; i < 4; i++ {
		fmt.Println("bar:", i)
	}
}
