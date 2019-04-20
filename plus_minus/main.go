// Given an array of integers, calculate the fractions of
// its elements that are positive, negative, and are zeros.
// Print the decimal value of each fraction on a new line.
//
// Example:
// A given array: [1, 1, 0, -1, -1] => 0.4 (2/5), 0.4 (2/5), 0.2 (1/5).
// Size of an array < 100.
// All elements in the array should be int32.
// Elements in array should have value < 100.
package main

import (
	"fmt"
)

type res struct {
	zeroRatio float64
	negRatio  float64
	posRatio  float64
}

var arr = []int32{1, 1, 0, -1, -1}

func plusMinus(arr []int32) res {
	var zeroCnt float64
	var posCnt float64
	var negCnt float64

	arrLen := float64(len(arr))
	for _, el := range arr {
		if el < 0 {
			negCnt++
		} else if el > 0 {
			posCnt++
		} else {
			zeroCnt++
		}
	}

	return res{
		zeroRatio: zeroCnt / arrLen,
		negRatio:  negCnt / arrLen,
		posRatio:  posCnt / arrLen,
	}
}

func main() {
	r := plusMinus(arr)
	fmt.Printf("%.6f\n", r.posRatio)
	fmt.Printf("%.6f\n", r.negRatio)
	fmt.Printf("%.6f\n", r.zeroRatio)
}
