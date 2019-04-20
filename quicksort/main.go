package main

import (
	"fmt"
)

var arr = []int{4, 10, 3, 9, 5} // 3,4,5,9,10

func quicksort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}
	pivot := unsorted[len(unsorted)-1]
	var left []int
	var right []int

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

func main() {
	sorted := quicksort(arr)
	fmt.Println(sorted)
}
