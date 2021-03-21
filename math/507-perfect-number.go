package main

func checkPerfectNumber(num int) bool {
	sum := -1 * num
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += num/i + i
			if i == num/i {
				sum -= i
			}
		}
	}
	return sum == num
}

func main() {
	checkPerfectNumber(28)
}
