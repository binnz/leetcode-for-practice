package main

func permute(nums []int) [][]int {
	res := [][]int{}
	var dfs func([]int, int, int)
	dfs = func(path []int, visited, dep int) {
		if dep == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for idx, e := range nums {
			if visited>>idx&1 == 1 {
				continue
			}
			path = append(path, e)
			visited ^= 1 << idx
			dfs(path, visited, dep+1)
			path = path[:len(path)-1]
			visited ^= 1 << idx
		}
	}
	dfs([]int{}, 0, 0)
	return res
}
