package main

import "fmt"

// #80 Remove Duplicates from Sorted Array II [Medium]
// Tags: Array, Two Pointers

func removeDuplicates(nums []int) int {

    left := 1

    for i := 2; i < len(nums); i++{
     if (nums[i] == nums[left] && nums[i] == nums[left-1]){
      continue
     }

     left ++
     nums[left] = nums[i]
    }
    
    left ++

    return left
}

func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,3})) // → 5
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,1,2,3,3})) // → 7
}
