package main

func findKthPositive(arr []int, k int) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		if arr[m]-(m+1) < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return k + l
}

// 鬼才解法
// 1. 缺失的正整数一定 >= k
// 2. 每当出现一个 <=k 的正整数， 少一个缺失的整数，k+1
/*
var findKthPositive = function (arr, k) {
    for (var i = 0; i < arr.length; i++) {
        if (arr[i] <= k) {
            k++
        }
    }

    return k
};
*/
