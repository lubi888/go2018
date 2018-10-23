package main

import (
	"encoding/json"
	"fmt"
)

type person8 struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person8{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person8{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person8{p1, p2}
	fmt.Println(people)
	fmt.Println(p2.Age)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err, "is the error found")
	}
	fmt.Println(bs)
	fmt.Println(string(bs), "json marshaled")

}
