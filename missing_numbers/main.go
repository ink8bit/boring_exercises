package main

import (
	"fmt"
)

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

func fillMap(arr []int32) map[int32]int32 {
	dict := map[int32]int32{}
	for _, el := range arr {
		_, ok := dict[el]
		if ok {
			dict[el]++
		} else {
			dict[el] = 1
		}
	}
	return dict
}

func missingNumbers(arr []int32, brr []int32) []int32 {
	dict := map[int32]int32{}
	res := []int32{}
	dict = fillMap(brr)
	for _, el := range arr {
		_, ok := dict[el]
		if ok {
			dict[el]--
		}
	}
	for k, v := range dict {
		if v > 0 {
			res = append(res, k)
		}
	}
	return quicksort(res)
}

func main() {
	brr := []int32{1, 2, 3, 5, 5, 5, 4}
	arr := []int32{1, 2, 3}
	fmt.Println(missingNumbers(arr, brr))
}
