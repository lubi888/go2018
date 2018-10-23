package main

import (
	"fmt"
	"runtime"
)

//var wg sync.WaitGroup
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

}
