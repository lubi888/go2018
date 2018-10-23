package main

import (
	"fmt"
	"sort"
)

type user12 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user12) String() string {
	return fmt.Sprintf("\n%s\t%s\t%d\n%s\n", u.First, u.Last, u.Age, u.Sayings)
}

type ByAge2 []user12

//func (a ByAge2) Less(i, j int) bool {
//	panic("implement me")
//}
//
//func (a ByAge2) Swap(i, j int) {
//	panic("implement me")
//}

func (a ByAge2) Len() int           { return len(a) }
func (a ByAge2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge2) Less(i, j int) bool { return a[i].Age < a[j].Age }

//func (bn ByAge2) Len() int           { return len(bn) }
//func (bn ByAge2) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
//func (bn ByAge2) Less(i, j int) bool { return bn[i].Age < bn[j].Age }

func main() {
	u1 := user12{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user12{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user12{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	user12s := []user12{u1, u2, u3}
	fmt.Println(user12s)
	fmt.Println()
	sort.Sort(ByAge2(user12s))
	fmt.Println(user12s)
}

//sort by name
//package main
//
//import (
//"fmt"
//"sort"
//)
//
//type Person struct {
//	First string
//	Age   int
//}
//
//type ByName []Person
//
//func (bn ByName) Len() int           { return len(bn) }
//func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
//func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }
//
//func main() {
//	p1 := Person{"James", 32}
//	p2 := Person{"Moneypenny", 27}
//	p3 := Person{"Q", 64}
//	p4 := Person{"M", 56}
//
//	people := []Person{p1, p2, p3, p4}
//
//	fmt.Println(people)
//	sort.Sort(ByName(people))
//	fmt.Println(people)
//
//}
