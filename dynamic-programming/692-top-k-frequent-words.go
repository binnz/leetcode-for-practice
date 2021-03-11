package main

import "sort"

func topKFrequent(words []string, k int) []string {
	wordCount := map[string]int{}
	countToWord := map[int][]string{}
	for _, e := range words {
		wordCount[e]++
	}
	for c, v := range wordCount {
		if _, ok := countToWord[v]; ok {
			countToWord[v] = append(countToWord[v], c)
		} else {
			countToWord[v] = []string{c}
		}
	}
	heap := []int{}
	for c := range countToWord {
		heap = append(heap, c)
	}
	res := []string{}
	heapLen := len(heap)
	buildHeap(heap, heapLen)
	for k > 0 && heapLen > 0 {
		words := countToWord[heap[0]]
		sort.Strings(words)
		idx := 0
		for k > 0 && idx < len(words) {
			res = append(res, words[idx])
			idx++
			k--
		}
		heapLen--
		heap[0], heap[heapLen] = heap[heapLen], heap[0]
		heapify(heap, heapLen, 0)

	}
	return res
}

func buildHeap(arr []int, n int) {
	for s := n/2 - 1; s >= 0; s-- {
		heapify(arr, n, s)
	}
}

func heapify(arr []int, n, idx int) {
	max := idx
	left := 2*idx + 1
	right := 2*idx + 2
	if left < n && arr[left] > arr[max] {
		max = left
	}
	if right < n && arr[right] > arr[max] {
		max = right
	}
	if max != idx {
		arr[max], arr[idx] = arr[idx], arr[max]
		heapify(arr, n, max)
	}
}
