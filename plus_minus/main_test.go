package main

import "testing"

var testCases = []struct {
	in  []int32
	out res
}{
	{
		[]int32{1, 1, 0, -1, -1},
		res{0.2, 0.4, 0.4},
	},
	{
		[]int32{1, -1, 0, 1, 0},
		res{0.4, 0.2, 0.4},
	},
}

func TestPlusMinus(t *testing.T) {
	for _, tt := range testCases {
		arr := tt.in
		want := tt.out
		got := plusMinus(tt.in)

		t.Run("", func(t *testing.T) {
			if got != want {
				t.Errorf("slice %v\ngot %v\nwant %v", arr, got, want)
			}
		})
	}
}
