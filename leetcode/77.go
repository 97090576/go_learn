package main

// cnk = cn-1k + cn-1k-1  所以这个题可以使用 DP 解，时间复杂度更低
func combine(n int, k int) [][]int {
	if 1 == k {
		var a [][]int
		for i := 1; i <= n; i++ {
			a = append(a, []int{i})
		}
		return a
	}
	if n == k {
		var a []int
		for i := 1; i <= n; i++ {
			a = append(a, i)
		}
		return [][]int{a}
	}
	res := combine(n-1, k-1)
	for i, v := range res {
		res[i] = append(v, n)
	}
	res = append(res, combine(n-1, k)...)
	return res
}
