package main

func updateMatrix(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	queue := []int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, i*col+j)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	way := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		i, j := cur/col, cur%col
		for k := 0; k < 4; k++ {
			x, y := i+way[k][0], j+way[k][1]
			if x >= 0 && x < row && y >= 0 && y < col && matrix[x][y] == -1 {
				matrix[x][y] = matrix[i][j] + 1
				queue = append(queue, x*col+y)
			}
		}
	}
	return matrix
}
