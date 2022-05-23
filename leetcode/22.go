package main

//
////var res []string
////
////func generateParenthesis(n int) []string {
////	backtrace(0, 0, n, "")
////	return res
////}
////
////func backtrace(l int, r int, n int, s string) {
////	if r == n || l == n {
////		if l == n {
////			for r < n {
////				s = s + ")"
////				r++
////			}
////		}
////		res = append(res, s)
////		return
////	}
////	if l == r {
////		backtrace(l+1, r, n, s+"(")
////	} else {
////		backtrace(l+1, r, n, s+"(")
////		backtrace(l, r+1, n, s+")")
////	}
////}
//
func generateParenthesis(n int) []string {
	var res []string
	s := ""
	var f func(l int, r int)
	f = func(l int, r int) {
		if r == 0 {
			res = append(res, s)
			return
		}
		if l == r {
			s += "("
			f(l-1, r)
			s = s[:len(s)-1]
		} else {
			if l > 0 {
				s += "("
				f(l-1, r)
				s = s[:len(s)-1]
			}
			s += ")"
			f(l, r-1)
			s = s[:len(s)-1]
		}
	}
	f(n, n)
	return res
}
