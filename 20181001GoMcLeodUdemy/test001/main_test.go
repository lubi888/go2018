package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3) //use sum to create error?
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
