package main

import "testing"

//func must be capital letter
func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		//unfurl
		x := mySum(v.data...)
		if x != v.answer {
			//must use t.Error
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
