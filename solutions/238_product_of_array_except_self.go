package main

import "fmt"

// #238 Product of Array Except Self [Medium]
// Tags: Array, Prefix Sum

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
  
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1,2,3,4})) // → [24,12,8,6]
	fmt.Println(productExceptSelf([]int{-1,1,0,-3,3})) // → [0,0,9,0,0]
}
