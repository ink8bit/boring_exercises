// Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.
//
// Note:
// Midnight is 12:00:00AM on a 12-hour clock,
// and 00:00:00 on a 24-hour clock.
// Noon is 12:00:00PM on a 12-hour clock, and 12:00:00 on a 24-hour clock.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var t = "07:05:45PM"

// formatHours returns formatted string of hours
// Example:
// "7" -> "07" (AM)
// "7" -> "19" (PM)
// "12" -> "00" (AM)
func formatHours(hours int, isAM bool) string {
	var hh string

	if (hours == 0 && isAM) || (hours < 10 && isAM) {
		hh = "0" + strconv.Itoa(hours)
	} else if hours == 12 && !isAM {
		hh = strconv.Itoa(hours)
	} else if hours == 12 && isAM {
		hours = 0
		hh = "0" + strconv.Itoa(hours)
	} else {
		hours = hours + 12
		hh = strconv.Itoa(hours)
	}

	return hh
}

// formatSeconds removes "PM" or "AM" suffix
// and returns a string containing only digits
func formatSeconds(ss string, isAM bool) string {
	if isAM {
		return strings.TrimSuffix(ss, "AM")
	}
	return strings.TrimSuffix(ss, "PM")
}

func timeConversion(s string) string {
	sl := strings.Split(s, ":")
	hh := sl[0]
	mm := sl[1]
	ss := sl[2]

	isAM := strings.HasSuffix(ss, "AM")

	hhInt, err := strconv.Atoi(hh)
	if err != nil {
		panic(err)
	}

	hh = formatHours(hhInt, isAM)
	ss = formatSeconds(ss, isAM)

	str := strings.Join([]string{hh, mm, ss}, ":")

	return str
}

func main() {
	s := timeConversion(t)
	fmt.Println(s)
}
