package main

import (
	"fmt"
	"math"
)

func main() {
	sq := square{
		length: float64(2.5),
	}
	fmt.Println(sq.area(), "is square area")

	ra := circle{
		radius: float64(20),
	}
	fmt.Println(ra.area(), "is the circle area")

	//anon func
	func() {
		fmt.Println("anon funco")
	}()
	info21(sq)
	info21(ra)
}
func info21(s shape) {
	fmt.Println(s.area())
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.length * s.length

}

type shape interface {
	area() float64
}

//package main
//
//import (
//"fmt"
//"math"
//)
//
//type circle struct {
//	radius float64
//}
//
//type square struct {
//	length float64
//}
//
//func (c circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}
//
//func (s square) area() float64 {
//	return s.length * s.length
//}
//
//type shape interface {
//	area() float64
//}
//
//func info(s shape) {
//	fmt.Println(s.area())
//}
//
//func main() {
//	circ := circle{
//		radius: 12.345,
//	}
//
//	squa := square{
//		length: 15,
//	}
//
//	info(circ)
//	info(squa)
//}
