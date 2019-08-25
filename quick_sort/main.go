package main

import (
	"fmt"
)

var arr = []int{4, 10, 3, 9, 5}

func quickSort(unsorted []int) []int {
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

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)
	sorted := append(sortedLeft, pivot)
	sorted = append(sorted, sortedRight...)

	return sorted
}

func main() {
	sorted := quickSort(arr)
	fmt.Println("unsorted:", arr)
	fmt.Println("sorted:", sorted)
}
