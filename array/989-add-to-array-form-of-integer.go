package main

func addToArrayForm(A []int, K int) []int {
	res := []int{}
	for i := len(A) - 1; i >= 0; i-- {
		val := A[i] + K%10
		K /= 10
		if val >= 10 {
			K++
			val -= 10
		}
		res = append(res, val)
	}
	for ; K > 0; K /= 10 {
		res = append(res, K%10)
	}
	for i, l := 0, len(res); i < len(res)/2; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	return res
}
