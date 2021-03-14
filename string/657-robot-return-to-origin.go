package main

func judgeCircle(moves string) bool {
	row := 0
	col := 0
	for _, e := range moves {
		if string(e) == "U" {
			row--
		} else if string(e) == "D" {
			row++
		} else if string(e) == "L" {
			col++
		} else {
			col--
		}
	}
	return row == 0 && col == 0
}
