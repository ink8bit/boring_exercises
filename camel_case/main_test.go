package main

import "testing"

var testCases = []struct {
	in   string
	out  int32
	name string
}{
	{
		in:   "",
		out:  0,
		name: "Empty string",
	},
	{
		in:   "oneTwoThree",
		out:  3,
		name: "oneTwoThree",
	},
}

func TestCamelCase(t *testing.T) {
	for _, tc := range testCases {
		want := tc.out
		got := camelCase(tc.in)

		t.Run(tc.name, func(t *testing.T) {
			if got != want {
				t.Errorf("got: %v, but want: %v", got, want)
			}
		})
	}
}
