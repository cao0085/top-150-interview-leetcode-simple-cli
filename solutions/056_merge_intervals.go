package main

import "fmt"

// #56 Merge Intervals [Medium]
// Tags: Array, Sorting

func merge(intervals [][]int) [][]int {

	i := 1
	for i < len(intervals) {
		c := intervals[i]
		p := intervals[i-1]
		if p[0] <= c[0] && c[0] <= p[1]{
			if p[1] < c[1] {
				p[1] = c[1]
			}
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			i++
		}
	}

	return intervals
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // → [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                   // → [[1,5]]
}
