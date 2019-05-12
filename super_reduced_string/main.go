package main

import (
	"fmt"
)

// superReducedString function returns "Empty String"
// if the given string is empty. If not the function
// selects a pair of adjacent lowercase letters that match,
// and deletes them.
// Example:
// "aaabccddd" -> "abd"
// "" -> "Empty String"
func superReducedString(s string) string {
	if len(s) == 0 {
		return "Empty String"
	}
	var res string
	dict := map[string]int{}
	for _, el := range s {
		elStr := string(el)
		_, ok := dict[elStr]
		if ok {
			dict[elStr]++
		} else {
			dict[elStr] = 1
		}
	}
	for k, v := range dict {
		if v%2 != 0 {
			res = res + k
		}
	}
	return res
}

func main() {
	str := "aabcc"
	fmt.Println(superReducedString(str))
}
