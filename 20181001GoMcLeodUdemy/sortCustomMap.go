package main

import (
	"fmt"
	"sort"
)

type person11 struct {
	Name string
	Age  int
}

func (p person11) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []person11 based on
// the Age field.
type ByAge []person11

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []person11{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
}



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
