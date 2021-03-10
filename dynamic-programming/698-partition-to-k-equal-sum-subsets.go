package main

func canPartitionKSubsets(nums []int, k int) bool {
	//[1,3,5,7,2,4,6,8]
	target, maxVal := 0, -2<<31
	for _, e := range nums {
		target += e
		if e < maxVal {
			maxVal = e
		}
	}
	if target%k != 0 {
		return false
	}
	target /= k
	if maxVal > target {
		return false
	}
	var dfs func(int, int, int, []bool) bool
	dfs = func(sum int, idx int, k int, used []bool) bool {
		if k == 0 {
			return true
		}
		if sum == target {
			return dfs(0, 0, k-1, used)
		}
		for i := idx; i < len(nums); i++ {
			if !used[i] && sum+nums[i] <= target {
				used[i] = true
				if dfs(sum+nums[i], i+1, k, used) {
					return true
				}
				used[i] = false
			}
		}
		return false
	}
	used := make([]bool, len(nums))
	return dfs(0, 0, k, used)
}
