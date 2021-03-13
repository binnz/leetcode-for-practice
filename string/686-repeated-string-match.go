package main

func repeatedStringMatch(a string, b string) int {
	var t string
	if len(a) < len(b) {
		for i := 0; i < len(b)/len(a)+2; i++ {
			t += a
		}
	} else {
		t = a + a
	}
	idx := matcher(t, b)
	if idx == -1 {
		return -1
	} else {
		if len(a) < len(b) {
			if (idx+len(b))%len(a) == 0 {
				return (idx + len(b)) / len(a)
			} else {
				return (idx+len(b))/len(a) + 1
			}
		} else {
			if idx+len(b) > len(a) {
				return 2
			} else {
				return 1
			}
		}
	}
	return -1
}

func matcher(a, pattern string) int {
	i, j := 0, 0
	lspArr := lsp(pattern)
	for j < len(a) {
		if a[j] == pattern[i] {
			i++
			j++
			if i == len(pattern) {
				return j - i
			}
		} else if i != 0 {
			i = lspArr[i-1]
		} else {
			j++
		}
	}
	return -1
}
func lsp(pattern string) []int {
	res := make([]int, len(pattern))
	pre, i := 0, 1
	for i < len(pattern) {
		if pre != 0 && pattern[i] != pattern[pre] {
			pre = res[pre-1]
		} else if pattern[i] == pattern[pre] {
			pre++
			res[i] = pre
			i++
		} else {
			i++
		}
	}
	return res
}
