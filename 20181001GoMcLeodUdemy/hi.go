package main

//2018.10.01     Udemy Todd McLeod
import (
	"fmt"
)

var y int = 2

type person struct {
	fname string //uppercase if want to be available outside package, not private
	lname string
	//Mname string
}

func (p person) speak() {
	fmt.Println("who are you?", p.fname, p.lname, `says "Goodmorn jo"`)
}

func (sa secretAgent) speak() {
	fmt.Println(`"shaken not stirred"`, sa.fname, sa.lname, `says "Goodmorn moneypenny"`)

}

type human interface { //for polymorphism
	speak()
}

func saySomething(h human) {
	h.speak()
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	x := 7
	fmt.Println(x)
	fmt.Printf("%T", x)
	fmt.Println()
	fmt.Println(y)
	xi := []int{2, 4, 6, 8}
	fmt.Println(xi)
	//for 2 := range xi {
	//	fmt.Println(xi)
	//}
	m := map[string]int{
		"beth": 44,
		"todd": 34,
	}
	fmt.Println(m)
	fmt.Println(m["todd"])

	n := map[int]string{
		44: "beth",
		34: "todd",
	}
	fmt.Println(n)
	fmt.Println(n[34])

	p1 := person{
		lname: string("moneypenny"), //otherwie just list in order without key and just value
		fname: string("miss"),
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"james",
			"bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()
	saySomething(p1)
	saySomething(sa1)
}
