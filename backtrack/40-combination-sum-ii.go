package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
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
			if sum+candidates[i] > target {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(path, sum+candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0, 0)
	return res
}
