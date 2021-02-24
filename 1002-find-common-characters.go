package main

func commonChars(A []string) []string {
	var (
		mid [26]int
		res [26]int
	)
	for i := 0; i < len(A[0]); i++ {
		res[A[0][i]-'a']++
	}
	for _, s := range A {
		mid = [26]int{}
		for i := 0; i < len(s); i++ {
			mid[s[i]-'a']++
		}
		for j := 0; j < 26; j++ {
			res[j] = min(res[j], mid[j])
		}
	}
	out := []string{}
	for i := 0; i < 26; i++ {
		for j := 0; j < res[i]; j++ {
			out = append(out, string('a'+i))
		}
	}
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
