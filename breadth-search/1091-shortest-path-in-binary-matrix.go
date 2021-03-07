package main

func shortestPathBinaryMatrix(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	way := [][]int{[]int{-1, -1}, []int{-1, 0}, []int{-1, 1}, []int{0, -1}, []int{0, 1}, []int{1, -1}, []int{1, 0}, []int{1, 1}}
	if grid[0][0] == 1 {
		return -1
	}
	if row == 1 && col == 1 {
		return 1
	}

	queue := []int{0}
	grid[0][0] = 1
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		i, j := cur/col, cur%col
		for k := 0; k < 8; k++ {
			x, y := i+way[k][0], j+way[k][1]
			if x >= 0 && x < row && y >= 0 && y < col && grid[x][y] == 0 {
				queue = append(queue, x*col+y)
				grid[x][y] = grid[i][j] + 1
				if x == row-1 && y == col-1 {
					return grid[x][y]
				}
			}
		}
	}
	return -1
}
