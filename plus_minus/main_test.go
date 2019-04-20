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
			t.Parallel()

			if got != want {
				t.Errorf("arr is %v and the result is %v, but want %v", arr, got, want)
			}
		})
	}
}
