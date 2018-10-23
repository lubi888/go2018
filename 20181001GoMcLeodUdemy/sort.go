package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}
	xx := make([]int, len(xi))
	xy := [9]int{}
	fmt.Println(xy, "is xy array")
	copy(xx, xi)
	fmt.Println(xi)
	fmt.Println(xs)
	fmt.Println(xx, "is xx copy")
	sort.Ints(xx)
	fmt.Println(xx, "xx is not sorted")
	fmt.Println("now begin sorting")
	sort.Ints(xi)
	fmt.Println(xi, "should sorted")
	sort.Strings(xs)
	fmt.Println(xs)
	for i, v := range xs {
		fmt.Println(i, v)
	}
	fmt.Println("are strings sorted ? assert")
	sort.StringsAreSorted(xs)
	fmt.Println("sorted strings check")
}
