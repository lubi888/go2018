package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkMyMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{2, 4, 6, 8})
	}
}

func ExampleCenteredAvg() {
	CenteredAvg([]int{1, 2, 3, 4, 5})
}

func TestCenteredAvg(t *testing.T) {
	n := CenteredAvg([]int{1, 3, 5})
	if n != 3 {
		fmt.Println("errors were found, sys")
		t.Error("got", n, "expected", 3)
	}
	//more tests
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{10, 20, 40, 60, 80}, 40},
		test{[]int{2, 4, 6, 8, 10, 12}, 7},
		test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Error("got", f, "want", v.answer)
		}
	}
}
