package main

func totalHammingDistance(nums []int) int {
	res := 0
	for i := 0; i <= 32; i++ {
		cOne := 0
		for _, e := range nums {
			if s := e & (1 << i); s != 0 {
				cOne++
			}
		}
		res += cOne * (len(nums) - cOne)
	}
	return res
}
