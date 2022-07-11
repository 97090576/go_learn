package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func d3() {
	s := ""
	fmt.Scan(&s)
	m := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	var a []int
	for i, v := range []byte(s) {
		if i != 2 {
			a = append(a, m[v])
		}
	}
	times := strings.Split(s, ":")
	hour := times[0]
	minute := times[1]
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(hour, minute)
	var hours []string
	var minutes []string
	hi, mi := 0, 0
	for _, v := range a {
		for j := 0; j < len(a); j++ {
			t := v*10 + a[j]
			ts := strconv.Itoa(t)
			if t < 10 {
				ts = "0" + "t"
			}
			if t < 24 {
				hours = append(hours, ts)
				if ts == hour {
					hi = len(hours)
				}
			}
			if t < 60 {
				minutes = append(minutes, ts)
				if ts == minute {
					mi = len(minutes)
				}
			}
		}
	}
	if mi == len(minutes) {
		fmt.Printf("%s:%s", hours[hi%len(hours)], minutes[0])
	} else {
		fmt.Printf("%s:%s", hour, minutes[mi])
	}
}
