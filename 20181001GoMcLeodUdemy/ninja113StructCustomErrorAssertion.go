package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")
	//	fmt.Println("foo ran -", e, "\n", e.(customErr).info)    //assertion
	//

}
//?
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
//	bs, err := toJSON(p1)
//
//	fmt.Println(string(bs))
//
//}
//
//// toJSON needs to return an error also
//func toJSON(a interface{}) []byte {
//	bs, _ := json.Marshal(a)
//}
