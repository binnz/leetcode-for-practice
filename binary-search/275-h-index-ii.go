package main

func hIndex1(citations []int) int {
	if len(citations) < 1 {
		return len(citations)
	}
	l, r := 0, len(citations)+1
	for l < r {
		m := l + (r-l)>>1
		left, right := 0, len(citations)
		for left < right {
			mid := left + (right-left)>>1
			if citations[mid] >= m {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if len(citations)-left >= m {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
