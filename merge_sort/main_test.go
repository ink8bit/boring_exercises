package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	in  []int
	out []int
}{
	{
		in:  []int{3, 2, 1, 4, 5},
		out: []int{1, 2, 3, 4, 5},
	},
	{
		in:  []int{17, 20, 110, 5, 55},
		out: []int{5, 17, 20, 55, 110},
	},
	{
		in:  []int{5},
		out: []int{5},
	},
}

func TestMergeSort(t *testing.T) {
	for _, tt := range testCases {
		arr := tt.in
		want := tt.out
		got := mergeSort(tt.in)

		t.Run("", func(t *testing.T) {
			t.Parallel()

			if !reflect.DeepEqual(got, want) {
				t.Errorf("❌ arr: is %v, got: %v, but wanted: %v", arr, got, want)
			}
		})
	}
}
