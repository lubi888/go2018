package main

import (
	"fmt"
	"sort"
)

type user54 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (p user54) String() string {
	return fmt.Sprintf("%s\t%s\t%d\n%s\n", p.First, p.Last, p.Age, p.Sayings)
}

type ByName []user54

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }


func main() {
	u1 := user54{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user54{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user54{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	user54s := []user54{u1, u2, u3}

	fmt.Println(user54s)
fmt.Println("000000000")
	sort.Sort(ByName(user54s))
	fmt.Println(user54s)

}





//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//type user54 struct {
//	first string
//	age   int
//}
//
//func (p user54) String() string {
//	return fmt.Sprintf("%s\t%d\n", p.first, p.age)
//}
//
//type ByName []user54
//
//func (bn ByName) Len() int           { return len(bn) }
//func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
//func (bn ByName) Less(i, j int) bool { return bn[i].first < bn[j].first }
//
//func main() {
//	p1 := user54{"James", 32}
//	p2 := user54{"Moneypenny", 27}
//	p3 := user54{"Q", 64}
//	p4 := user54{"M", 56}
//
//	people3 := []user54{p1, p2, p3, p4}
//
//	fmt.Println(people3)
//	sort.Sort(ByName(people3))
//	fmt.Println(people3)
//}
//
////func byName(user54s []user54) sort.Interface {
////	func (a ByAge) Len() int           { return len(a) }
////	func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
////	func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
////}
