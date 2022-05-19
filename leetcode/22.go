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
//func generateParenthesis(n int) []string {
//	res := []string{"("}
//	l, r := 1, 0
//	for r < n {
//		if l == r {
//			for _, v := range res {
//				res = append(res, v+"(")
//				res = res[1:]
//			}
//			l++
//		} else {
//			for _, v := range res {
//				res = append(res, v+"(")
//				res = res[1:]
//			}
//			l++
//		}
//	}
//
//}
