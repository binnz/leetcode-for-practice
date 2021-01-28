package main

func balancedString(s string) int {
	left, right, res := 0, 0, len(s)
	k, counter := len(s)/4, make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}
	for left <= right && right<=len(s) {
		if counter['Q'] > k || counter['W'] > k || counter['E'] > k || counter['R'] > k{
			if right<len(s){
				right++
				counter[s[right-1]]--
			}else{
				break
			}
		} else {
				res = min(res, right-left)
				counter[s[left]]++
				left++
		}

	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main(){
	s:="QWER"
	balancedString(s)
}