package dog

import (
	"fmt"
	"testing"
)

//type bench

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		fmt.Println("there has been an error", t)
		t.Error("got", n, "wanted 70", 70)
	}
}
func TestYearsTwo(t *testing.T) {
	n := YearsTwo(5)
	if n != 35 {
		fmt.Println("error again, expected 35")
		t.Error("got", n, "expected ", 37)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//output 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

//type test
//func TestYears(t *testing.T) {
//	if Years(5) != 35{
//		//return fmt.Sprintln("errors:", t)
//	}
//}
