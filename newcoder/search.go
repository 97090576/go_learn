package main

func search(paths []string, qx map[string]byte) map[string]byte {
	res := make(map[string]byte)
	for _, v := range paths {
		res[v] = backtrace(v, qx)
	}
	return res
}

func backtrace(path string, qx map[string]byte) byte {
	b, f := qx[path]
	if f == true {
		return b
	} else {
		i := last(path)
		if i == 0 {
			return 'N'
		}
		return backtrace(path[:i], qx)
	}
}

func last(s string) int {
	i := len(s)
	for i > 0 {
		i--
		if s[i] == '/' {
			return i
		}
	}
	return 0
}
