//main package & work of package commenting
package main

//importets needed
import "fmt"
import "./dog"

//Canine structure with name and age
type canine struct {
	name string
	age  int
}

//unused other func
//func (c canine) YearsD(h int) int  {
//	return c.age * 7
//}

//main func entry point
func main() {
	fmt.Println("dog package prog")

	var h int
	fmt.Println("here is the dog years exported func")
	fmt.Scan(&h)
	dy := dog.Years(h)
	fmt.Println("the dow is now ", dy, "years in human")

	c1 := canine{
		name: "bobby",
		age:  dog.Years(h),
	}
	fmt.Println("the canine is now human years", dy, "and his name is", c1.name)
}
