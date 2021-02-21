package main

import "fmt"

func main() {
	hm := make(map[int]int)
	v, ok := hm[0]
	fmt.Println(v, ok)
	hm[0]++
	t, o := hm[0]
	fmt.Println(t, o)
}
