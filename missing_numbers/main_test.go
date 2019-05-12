package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	in   []int32
	orig []int32
	out  []int32
}{
	{
		in:   []int32{7, 2, 5, 3, 5, 3},
		orig: []int32{7, 2, 5, 4, 6, 3, 5, 3},
		out:  []int32{4, 6},
	},
	{
		in:   []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206},
		orig: []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204},
		out:  []int32{204, 205, 206},
	},
	{
		in:   []int32{1, 2, 3},
		orig: []int32{1, 2, 3, 5, 5, 5, 4},
		out:  []int32{4, 5},
	},
}

func TestMissingNumbers(t *testing.T) {
	for _, tc := range testCases {
		in := tc.in
		orig := tc.orig
		want := tc.out
		got := missingNumbers(in, orig)

		t.Run("", func(t *testing.T) {
			t.Parallel()

			if !reflect.DeepEqual(got, want) {
				t.Errorf("arr: is %v, got: %v, but wanted: %v", in, got, want)
			}
		})
	}
}
