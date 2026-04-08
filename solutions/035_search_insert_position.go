package main

import "fmt"

// #35 Search Insert Position [Easy]
// Tags: Array, Binary Search

func searchInsert(nums []int, target int) int {
    
    for i, v := range nums {

      if v == target {
        return i
      }

      if v > target {
        return i
      }

    }
    return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5)) // → 2
	fmt.Println(searchInsert([]int{1,3,5,6}, 2)) // → 1
	fmt.Println(searchInsert([]int{1,3,5,6}, 7)) // → 4
	fmt.Println(searchInsert([]int{1,3,5,6}, 0)) // → 0
}
