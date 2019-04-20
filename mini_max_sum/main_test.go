package main

import (
	"testing"
)

var sortedTestCases = []struct {
	arr []int32
	min int64
	max int64
}{
	{
		arr: []int32{1, 2, 3, 4, 5},
		min: 10,
		max: 14,
	},
	{
		arr: []int32{1, 3, 5, 7, 9},
		min: 16,
		max: 24,
	},
}

var unsortedTestCases = []struct {
	arr []int32
	min int64
	max int64
}{
	{
		arr: []int32{1, 2, 3, 4, 5},
		min: 10,
		max: 14,
	},
	{
		arr: []int32{1, 3, 5, 7, 9},
		min: 16,
		max: 24,
	},
	{
		arr: []int32{7, 69, 2, 221, 8974},
		min: 299,
		max: 9271,
	},
}

func TestMiniMaxSum(t *testing.T) {
	// t.Skip("skipping test")
	for _, tt := range sortedTestCases {
		arr := tt.arr
		wantMin := tt.min
		wantMax := tt.max
		gotMin, gotMax := miniMaxSum(arr)

		t.Run("", func(t *testing.T) {
			t.Parallel()
			if gotMin != wantMin || gotMax != wantMax {
				t.Errorf(
					"arr: %v, min sum %v, max sum %v,\nbut wanted min sum %v and max sum %v",
					arr, gotMin, gotMax, wantMin, wantMax,
				)
			}
		})
	}
}

func TestUMiniMaxSum(t *testing.T) {
	for _, tt := range unsortedTestCases {
		arr := tt.arr
		wantMin := tt.min
		wantMax := tt.max
		gotMin, gotMax := uMiniMaxSum(arr)

		t.Run("", func(t *testing.T) {
			t.Parallel()
			if gotMin != wantMin || gotMax != wantMax {
				t.Errorf(
					"arr: %v, min sum %v, max sum %v,\nbut wanted min sum %v and max sum %v",
					arr, gotMin, gotMax, wantMin, wantMax,
				)
			}
		})
	}
}
