package main

func maxProduct(words []string) int {
	bitRep := make([]int, len(words))
	for idx, e := range words {
		num := 0
		for _, i := range e {
			num |= 1 << (i - 97)
		}
		bitRep[idx] = num
	}
	res := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if bitRep[i]&bitRep[j] == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
