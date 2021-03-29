package main

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	var dfs func([]int, int, int, int)
	dfs = func(path []int, sum, dep, start int) {
		if sum == n && dep == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < 10; i++ {
			path = append(path, i)
			dfs(path, sum+i, dep+1, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0, 0, 1)
	return res
}
