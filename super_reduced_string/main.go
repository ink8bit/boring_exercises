package main

import (
	"fmt"
	"strings"
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

	res := ""
	for _, el := range s {
		elStr := string(el)
		if !strings.Contains(res, elStr) && dict[elStr]%2 != 0 {
			res += elStr
		} else {
			continue
		}
	}

	if res == "" {
		return "Empty String"
	}

	return res
}

func main() {
	str := "aaabccddd"
	fmt.Println(superReducedString(str))
}
