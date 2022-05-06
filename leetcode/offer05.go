package main

func replaceSpace(s string) string {
	ss := []byte(s)
	var count int
	for _, v := range ss {
		if v == ' ' {
			count++
		}
	}
	i := 0
	res := make([]byte, len(s)+2*count)
	for _, v := range ss {
		if v != ' ' {
			res[i] = v
		} else {
			res[i] = '%'
			i++
			res[i] = '2'
			i++
			res[i] = '0'
		}
		i++
	}
	return string(res)
}
