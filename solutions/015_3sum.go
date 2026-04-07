package main

import "fmt"
import "sort"

// #15 3Sum [Medium]
// Tags: Array, Two Pointers, Sorting

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    left := 0
    right := len(nums) - 1
    result := [][]int{}

    for left < right {

      n1 := nums[left]
      var n2 int
      n3 := nums[right]
      
      // fmt.Println(n1,n3)
      if n1 + n3 > 0 { // 太大
        right --
        n2 = nums[right]
      } else {
        left ++
        n2 = nums[left]
      }

      if n1 + n2 + n3 == 0{
        result = append(result,[]int{n1,n2,n3})
      }
    }
    return result
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4})) // → [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0,0,0})) // → [[0,0,0]]
	fmt.Println(threeSum([]int0,1,1})) // → []
}
