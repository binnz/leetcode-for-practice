package main

func canReach(arr []int, start int) bool {
	var dfs func(int)
	var res bool
	l := len(arr)
	dfs = func(n int) {
		if n >= l || n < 0 || arr[n] == -1 {
			return
		}
		if res {
			return
		}
		step := 0
		if arr[n] == 0 {
			res = true
			return
		} else {
			step = arr[n]
			arr[n] = -1
		}
		dfs(n - step)
		dfs(n + step)

	}
	dfs(start)
	return res
}
