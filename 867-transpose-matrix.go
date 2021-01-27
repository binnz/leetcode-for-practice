package main

func transpose(A [][]int) [][]int {
	row, col := len(A), len(A[0])
	res := make([][]int, col)
	for i := 0; i < col; i++ {
		res[i] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res[j][i] = A[i][j]
		}
	}
	return res
}
