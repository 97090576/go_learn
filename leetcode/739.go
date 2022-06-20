package main

func dailyTemperatures(temperatures []int) (res []int) {
	var s []p
	for i := len(temperatures); i > 0; {
		i--
		v := temperatures[i]
		if i == len(temperatures)-1 {
			s = append(s, p{i, v})
			res = append(res, 0)
		} else {
			for len(s) > 0 && s[len(s)-1].v < v {
				s = s[:len(s)-1]
			}
			if len(s) == 0 {
				res = append(res, 0)
			} else {
				res = append(res, s[len(s)-1].i-i)
			}
			s = append(s, p{i, v})
		}
	}
	reverse(res, 0, len(res))
	return
}

type p struct {
	i, v int
}
