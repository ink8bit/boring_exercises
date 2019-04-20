// Given five positive integers, find the minimum and maximum values
// that can be calculated by summing exactly four of the five integers.
// Then print the respective minimum and maximum values as a single line
// of two space-separated long integers.
//
// Example:
// Gievn an array [1, 3, 5, 7, 9].
// Our minimum sum is 1 + 3 + 5 + 7 = 16 and
// our maximum sum is 3 + 5 + 7 + 9 = 24.
// We would print 16 24.
//
// Note: The output can be greater than a 32 bit integer.
package main

import (
	"fmt"
	_ "sort"
)

var sorted = []int32{1, 2, 3, 4, 5}
var unsorted = []int32{7, 69, 2, 221, 8974} // 299 9271

// miniMaxSum returns min and max sums of elements
// in a given SORTED array
func miniMaxSum(sortedArr []int32) (int64, int64) {
	var sum int64

	for _, el := range sortedArr {
		sum += int64(el)
	}

	min := sum - int64(sortedArr[len(sortedArr)-1])
	max := sum - int64(sortedArr[0])

	return min, max
}

func quicksort(unsorted []int32) []int32 {
	if len(unsorted) < 2 {
		return unsorted
	}
	pivot := unsorted[len(unsorted)-1]
	var left []int32
	var right []int32

	for _, el := range unsorted[0 : len(unsorted)-1] {
		if el < pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	sortedLeft := quicksort(left)
	sortedRight := quicksort(right)
	sorted := append(sortedLeft, pivot)
	sorted = append(sorted, sortedRight...)

	return sorted
}

// miniMaxSum returns min and max sums of elements
// in a given UNSORTED array
func uMiniMaxSum(arr []int32) (int64, int64) {
	var sum int64

	sorted := quicksort(arr)

	for _, el := range sorted {
		sum += int64(el)
	}

	min := sum - int64(sorted[len(sorted)-1])
	max := sum - int64(sorted[0])

	return min, max
}

func main() {
	min, max := miniMaxSum(sorted)
	fmt.Println(min, max)

	umin, umax := uMiniMaxSum(unsorted)
	fmt.Println(umin, umax)
}
