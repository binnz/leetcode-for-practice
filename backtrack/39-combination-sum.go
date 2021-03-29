package main

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func([]int, int, int)
	dfs = func(path []int, sum int, start int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(path, sum+candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0, 0)
	return res
}
