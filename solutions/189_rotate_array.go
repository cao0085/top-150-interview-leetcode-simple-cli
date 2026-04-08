package main

import "fmt"

// #189 Rotate Array [Medium]
// Tags: Array, Math, Two Pointers

func rotate(nums []int, k int) {
 n := len(nums) - k
 reverse(nums)
 reverse(nums[:n])
 reverse(nums[n:])
 fmt.Println(nums,k)
}

func reverse(nums []int){
  
 left := 0
 right := len(nums)-1
 for left < right {
  nums[left], nums[right] = nums[right], nums[left]

  left ++
  right --

 }
}


func main() {
	nums1 := []int{1,2,3,4,5,6,7}
	rotate(nums1, 3)
	fmt.Println(nums1) // → [5,6,7,1,2,3,4]

	nums2 := []int{-1,-100,3,99}
	rotate(nums2, 2)
	fmt.Println(nums2) // → [3,99,-1,-100]
}
