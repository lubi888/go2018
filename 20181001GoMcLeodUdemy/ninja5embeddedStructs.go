package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type sedan struct {
	vehicle
	luxury bool
}

//moved to anon strct in main()
//type truck struct {
//	vehicle
//	fourWheel bool
//}

func main() {

	s1 := sedan{
		vehicle: vehicle{
			doors:  int(4),
			colour: string("purple"),
		},
		luxury: true,
	}

	s2 := struct {
		vehicle
		fourWheel bool
	}{
		fourWheel: true,
		vehicle: vehicle{
			doors:  5,
			colour: string("yellow"),
		},
	}

	fmt.Println(s1.colour, s1.luxury)
	fmt.Println(s2.colour, s2.doors)
}
