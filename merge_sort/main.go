package main

import (
	"fmt"
)

var arr = []int{3, 2, 1, 4, 5}

// merge function concatenates two given arrays
// into a sorted one.
func merge(left, right []int) []int {
	var sorted []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}
	sorted = append(sorted, left...)
	sorted = append(sorted, right...)
	return sorted
}

// mergeSort function divides a given array into
// two subarrays and merges them into a sorted array.
func mergeSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}
	middle := len(unsorted) / 2
	left := unsorted[0:middle]
	right := unsorted[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func main() {
	fmt.Println("unsorted:", arr)
	fmt.Println("sorted:", mergeSort(arr))
}
