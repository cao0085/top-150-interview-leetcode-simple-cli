package main

import "fmt"

// #88 Merge Sorted Array [Easy]
// Tags: Array, Two Pointers, Sorting

func merge(nums1 []int, m int, nums2 []int, n int) {
    
}

func main() {
	fmt.Println(merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)) // → [1,2,2,3,5,6]
	fmt.Println(merge([]int{1}, 1, []int{}, 0)) // → [1]
}
