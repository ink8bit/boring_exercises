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
	var sl []rune
	for _, ch := range s {
		if len(sl) == 0 {
			sl = append(sl, ch)
		} else {
			if sl[len(sl)-1] == ch {
				sl = sl[:len(sl)-1]
			} else {
				sl = append(sl, ch)
			}
		}
	}
	res := string(sl)
	if res == "" {
		return "Empty String"
	}
	return res
}

func main() {
	str := "aaabbccddd"
	fmt.Println(superReducedString(str))
}
