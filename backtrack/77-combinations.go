package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	var dfs func([]int, int, int)
	dfs = func(path []int, start int, dep int) {
		if dep == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			dfs(path, i+1, dep+1)
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 1, 0)
	return res
}
