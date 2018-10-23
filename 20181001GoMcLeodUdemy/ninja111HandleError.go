package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error happened here: ", err)
		log.Fatalln("well that didn't work", err)
		log.Println("error log here: ", err)
		panic(err)
	}
	fmt.Println(string(bs))

}

//given code

//package main
//
//import (
//"encoding/json"
//"fmt"
//)
//
//type person struct {
//	First   string
//	Last    string
//	Sayings []string
//}
//
//func main() {
//	p1 := person{
//		First:   "James",
//		Last:    "Bond",
//		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
//	}
//
//	bs, _ := json.Marshal(p1)
//	fmt.Println(string(bs))
//
//}
