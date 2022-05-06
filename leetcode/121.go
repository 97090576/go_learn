package main

func maxProfit(prices []int) int {
	res := 0
	low := prices[0]
	for _, v := range prices {
		if v-low > res {
			res = v - low
		}
		if v < low {
			low = v
		}
	}
	return res
}
