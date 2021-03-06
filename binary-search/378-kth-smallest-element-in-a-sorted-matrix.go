package main

func kthSmallest(matrix [][]int, k int) int {
	left, right := matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]
	for left < right {
		mid := left + (right-left)>>1
		count := countLessOrEqualThanTarget(matrix, mid)
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func countLessOrEqualThanTarget(matrix [][]int, t int) int {
	i, j := len(matrix)-1, 0
	c := 0
	for j < len(matrix[0]) && i >= 0 {
		if matrix[i][j] <= t {
			c += i + 1
			j++
		} else {
			i--
		}
	}
	return c
}

func main() {
	kthSmallest([][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 7)
}
