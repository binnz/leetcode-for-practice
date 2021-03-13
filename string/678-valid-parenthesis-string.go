package main

func checkValidString(s string) bool {
	left, star := []int{}, []int{}
	for i, e := range s {
		if string(e) == "(" {
			left = append(left, i)
		} else if string(e) == "*" {
			star = append(star, i)
		} else {
			if len(left) == 0 && len(star) == 0 {
				return false
			} else if len(left) != 0 {
				left = left[:len(left)-1]
			} else {
				star = star[:len(star)-1]
			}
		}
	}
	for len(star) != 0 && len(left) != 0 {
		if star[len(star)-1] < left[len(left)-1] {
			return false
		}
		left = left[:len(left)-1]
		star = star[:len(star)-1]
	}
	if len(left) != 0 {
		return false
	}
	return true
}
