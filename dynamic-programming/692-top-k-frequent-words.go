package main

func topKFrequent(words []string, k int) []string {
	wordCount := map[string]int{}
	countToWord := map[int][]string{}
	for _, e := range words {
		wordCount[e]++
	}
	for k, v := range wordCount {
		if _, ok := countToWord[v]; ok {
			countToWord[v] = append(countToWord[v], k)
		} else {
			countToWord[v] = []string{k}
		}
	}
	heap := []int{}

}
func buildHeap([]int) {}

func heapify()
