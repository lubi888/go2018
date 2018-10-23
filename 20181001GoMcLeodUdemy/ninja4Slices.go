package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Print(a, "\n")

	for i, v := range a {
		fmt.Println(i, v)
	}
	b := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Print(b, "\n")
	for i, v := range b {
		fmt.Print(i, v, "\t")
		fmt.Printf("%T", v)
		fmt.Println()
	}

	y1 := append(b[:2], b[4:]...)
	fmt.Println(b)
	for i, ok := range y1 {
		fmt.Println(i, ok)
	}

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Print(x, "\t")
	fmt.Println()
	x = append(x, 53, 54, 55)
	fmt.Print(x, "\t")
	fmt.Println()
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	z := append(x[:2], x[5:]...)
	fmt.Print(z, "\t")

	usa := []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}

	fmt.Println("\n", usa, "\t")
	fmt.Print(len(usa), cap(usa), "\n") //
	for g, ok := range usa {
		fmt.Print(g, ok, "\n")
	}

	j := make([]string, 4, 4)
	j = []string{"James", "Bond", "Shaken, not stirred"}
	b1 := make([]string, 4, 4)
	b1 = []string{"Miss", "Moneypenny", "Helloooooo, James."}

	f := [][]string{j, b1}
	for i, v := range f {
		fmt.Print(i, v, "\n")
		for u, v := range f {
			fmt.Printf("\t index position: %v \t value: \t %v \n", u, v)
		}
	}

	m := map[string][]string{
		//"last_firstname": "value",
		"bond_james":       {"Shaken not stirred", "martinis", "women"},
		"no_dr":            {"being evil", "island popping", "toys"},
		"moneypennhy_miss": {"daquiris", "makeup", "networking"},
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k)
		for l, u := range v {
			fmt.Println("\t", l, u)
		}
	}

	m["q"] = []string{"umbrella daggers", "gas bullets", "aligators airtanks"}
	for k, v := range m {
		fmt.Println(k)
		for l, u := range v {
			fmt.Println("\t", l, u)
		}
	}

	delete(m, "no_dr")
	fmt.Println(len(m))
	for k, v := range m {
		fmt.Println(k)
		for l, u := range v {
			fmt.Println("\t", l, u)
		}
	}

}
