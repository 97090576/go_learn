package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], v)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
