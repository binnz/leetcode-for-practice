package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	row := len(grid)
	col := len(grid[0])
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		s := 1
		s += dfs(i, j-1)
		s += dfs(i, j+1)
		s += dfs(i-1, j)
		s += dfs(i+1, j)
		return s
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
