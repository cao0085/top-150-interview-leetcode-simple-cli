package main

import "fmt"

// #274 H-Index [Medium]
// Tags: Array, Sorting, Counting Sort

func hIndex(citations []int) int {
	
	n := len(citations)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if citations[j] < citations[j+1] {
				citations[j], citations[j+1] = citations[j+1], citations[j]
			}
		}
	}

	for i, v := range citations {
		if v < i+1 {
			return i
		}
	}
	return n
}

func main() {
	fmt.Println(hIndex([]int{3,0,6,1,5})) // → 3
	fmt.Println(hIndex([]int{1,3,1})) // → 1
}
