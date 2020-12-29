package main

import (
	"testing"
)

var testCases = []struct {
	in  string
	out string
}{
	{
		in: "12:00:00AM", out: "00:00:00",
	},
	{
		in: "12:00:00PM", out: "12:00:00",
	},
	{
		in: "07:05:45PM", out: "19:05:45",
	},
	{
		in: "05:15:50PM", out: "17:15:50",
	},
	{
		in: "07:05:45AM", out: "07:05:45",
	},
	{
		in: "05:15:50AM", out: "05:15:50",
	},
}

func TestTimeConversion(t *testing.T) {
	// t.Skip("time conversion")

	for _, tt := range testCases {
		str := tt.in
		want := tt.out
		got := timeConversion(tt.in)

		t.Run("", func(t *testing.T) {
			if got != want {
				t.Errorf("string %q\ngot %q\nwant %q", str, got, want)
			}
		})
	}
}
