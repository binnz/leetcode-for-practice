package main

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	for j := len(T) - 2; j >= 0; j-- {
		if T[j] < T[j+1] {
			res[j] = 1
		} else {
			if res[j+1] == 0 {
				res[j] = 0
			} else {
				i := j + 1 + res[j+1]
				for i < len(T)-1 {
					if T[j] < T[i] {
						res[j] = i - j
						break
					} else if res[i] > 0 {
						i = i + res[i]
					} else {
						break
					}
				}

			}
		}
	}
	return res
}

func main() {
	dailyTemperatures([]int{77, 77, 77, 77, 77, 41, 77, 41, 41, 77})
}
