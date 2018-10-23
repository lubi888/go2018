package main

import (
	"encoding/json"
	"fmt"
)

type person9 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people5 := []person9{}
	var people6 []person9
	fmt.Println(people6, "is people 6slice")

	err := json.Unmarshal(bs, &people6)
	if err != nil {
		fmt.Println(err, "is the rrrror")
	}

	fmt.Println("all the data peoples raw", people6)

	for i, v := range people6 {
		fmt.Println("people number", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		fmt.Println()
	}
}
