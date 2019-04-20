package main

import (
	"fmt"
)

var arr = []int{3, 2, 1, 4, 5}

func merge(left, right []int) []int {
	var sorted []int

	// todo
	return sorted
}

func mergeSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}

	pivot := len(unsorted) / 2 // 2
	left := unsorted[0:pivot]
	right := unsorted[pivot:]
	fmt.Println(left)
	fmt.Println(right)
	sortedLeft := mergeSort(left)
	sortedRight := mergeSort(right)

	return merge(sortedLeft, sortedRight)
}

func main() {
	fmt.Println("unsorted array:", arr)
	mergeSort(arr)
}
