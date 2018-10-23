package word

import (
	"fmt"
	"testing"
	"../quote"
	)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	//should be 1349
}
func ExampleUseCount() {
	Count(quote.SunAlso)
}

func TestCount(t *testing.T) {
	l := Count(quote.SunAlso)
	if l != 1349 {
		t.Error("got", l, "expected", 1349)
	}
}

func TestUseCount(t *testing.T) {
	l := UseCount("one two two three")
	for k, v := range l {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "two":
			if v != 2 {
				t.Error("got", v, "exptd", 2)
			}
		case "three":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		}
	}
}
