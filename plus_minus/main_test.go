package main

import (
	"testing"
)

func TestPlusMinus(t *testing.T) {
	arr := []int32{1, 1, 0, -1, -1}
	want := res{0.2, 0.4, 0.4}
	got := plusMinus(arr)

	if got != want {
		t.Errorf("arr is %v and the result is %v, but want %v", arr, got, want)
	}
}
