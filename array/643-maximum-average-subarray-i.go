package main

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxV := float64(sum) / float64(k)
	for i := 1; i <= len(nums)-k; i++ {
		sum -= nums[i-1]
		sum += nums[i+k-1]
		if float64(sum)/float64(k) > maxV {
			maxV = float64(sum) / float64(k)
		}
	}
	return maxV
}
