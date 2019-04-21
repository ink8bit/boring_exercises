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

	if hhInt == 0 && isAM {
		hh = "0" + strconv.Itoa(hhInt)
	} else if hhInt < 10 && isAM {
		hh = "0" + strconv.Itoa(hhInt)
	} else if hhInt == 12 && !isAM {
		hhInt = 12
		hh = strconv.Itoa(hhInt)
	} else if hhInt == 12 && isAM {
		hhInt = 0
		hh = "0" + strconv.Itoa(hhInt)
	} else {
		hhInt = hhInt + 12
		hh = strconv.Itoa(hhInt)
	}

	if isAM {
		ss = strings.TrimSuffix(ss, "AM")
	} else {
		ss = strings.TrimSuffix(ss, "PM")
	}

	str := strings.Join([]string{hh, mm, ss}, ":")

	return str
}

func main() {
	s := timeConversion(t)
	fmt.Println(s)
}
