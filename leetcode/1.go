package main

func twoSum_1(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{m[target-v], i}
		}
		m[v] = i
	}
	return nil
}
