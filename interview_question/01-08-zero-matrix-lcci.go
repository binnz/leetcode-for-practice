package main

func setZeroes(matrix [][]int) {
	rhm := map[int]struct{}{}
	chm := map[int]struct{}{}
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		if _, ok := rhm[i]; ok {
			continue
		}
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rhm[i] = struct{}{}
				chm[j] = struct{}{}
			}
		}
	}
	for k, _ := range rhm {
		for j := 0; j < col; j++ {
			matrix[k][j] = 0
		}
	}
	for k, _ := range chm {
		for i := 0; i < row; i++ {
			matrix[i][k] = 0
		}
	}
}
