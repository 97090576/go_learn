package main

import (
	"sort"
)

// 首先二分计数每一行的 1 的数量，保留数量和下标索引的映射，然后取最弱的 K 个 （理论上这个应该使用一个最小堆）， K 个节点的最小堆
// 二分查找 + 最小堆
func kWeakestRows(mat [][]int, k int) []int {
	s := make([]pair, 0)
	for j, v := range mat {
		i := sort.Search(len(v), func(i int) bool { return v[i] < 1 })
		s = append(s, pair{i, j})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].x < s[j].x || (s[i].x == s[j].x && s[i].y < s[j].y)
	})
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, s[i].y)
	}
	return res
}

/* 官方 手动实现了最小堆，说明 go 标准库确实是没有堆的
func kWeakestRows(mat [][]int, k int) []int {
    h := hp{}
    for i, row := range mat {
        pow := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
        h = append(h, pair{pow, i})
    }
    heap.Init(&h)
    ans := make([]int, k)
    for i := range ans {
        ans[i] = heap.Pop(&h).(pair).idx
    }
    return ans
}

type pair struct{ pow, idx int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.pow < b.pow || a.pow == b.pow && a.idx < b.idx }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
*/
