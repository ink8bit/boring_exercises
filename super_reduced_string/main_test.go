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
		in:   "",
		out:  "Empty String",
		name: "An empty string",
	},
	{
		in:   "aaabccddd",
		out:  "abd",
		name: "aaabccddd",
	},
}

func TestSuperReducedString(t *testing.T) {
	for _, tc := range testCases {
		want := tc.out
		got := superReducedString(tc.in)

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got != want {
				t.Errorf("got: %v, but want: %v", got, want)
			}
		})
	}
}
