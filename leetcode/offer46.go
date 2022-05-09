package main

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	// 1. 把 num 拆成数组
	var r []int
	for num > 0 {
		r = append(r, num%10)
		num /= 10
	}
	reverse(r, 0, len(r))
	f := 1
	s := 1
	if r[0] != 0 && r[0]*10+r[1] <= 25 {
		s++
	}
	for i := 2; i < len(r); i++ {
		tmp := s
		if r[i-1] != 0 && r[i-1]*10+r[i] <= 25 {
			tmp += f
		}
		f = s
		s = tmp
	}
	return s
}
