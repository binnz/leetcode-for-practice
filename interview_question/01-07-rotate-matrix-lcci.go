package main

func rotate(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := i + 1; j < col; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		for j := 0; j < col/2; j++ {
			matrix[i][j], matrix[i][col-j-1] = matrix[i][col-j-1], matrix[i][j]
		}
	}
}
