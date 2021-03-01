package main

func maxSubArray(nums []int) int {
	res := -101
	t := 0
	for _, e := range nums {
		if t < 0 {
			t = e
		} else {
			t += e
		}
		res = max(res, t)
	}
	return res

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
