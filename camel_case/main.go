package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "oneTwoThree"
	fmt.Println(camelcase(str))
}

// camelcase function returns a number of
// words in camel-cased string
//
// Example:
// Input: oneTwoThree -> Output: 3
func camelcase(s string) int32 {
	if len(s) == 0 {
		return 0
	}

	var cnt int32
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			cnt++
		}
	}
	return cnt + 1
}
