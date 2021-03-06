package main

import (
	"testing"
)

var testCases = []struct {
	in   string
	out  string
	name string
}{
	{
		in:   "aabcc",
		out:  "b",
		name: "aabcc",
	},
	{
		in:   "aacc",
		out:  "Empty String",
		name: "aacc",
	},
	{
		in:   "",
		out:  "Empty String",
		name: "An empty string",
	},
	{
		in:   "aaabccddd",
		out:  "abd",
		name: "aaabccddd",
	},
	{
		in:   "aa",
		out:  "Empty String",
		name: "aa",
	},
	{
		in:   "baab",
		out:  "Empty String",
		name: "baab",
	},
	{
		in:   "aaaaabbccddd",
		out:  "ad",
		name: "aaaaabbccddd",
	},
}

func TestSuperReducedString(t *testing.T) {
	for _, tc := range testCases {
		want := tc.out
		got := superReducedString(tc.in)

		t.Run(tc.name, func(t *testing.T) {
			if got != want {
				t.Errorf("got %v\nwant %v", got, want)
			}
		})
	}
}
