package main

import "fmt"

// #1 Two Sum [Easy]
// Tags: Array, Hash Table

func twoSum(nums []int, target int) []int {
    needSet := make(map[int]int)
    for i, n := range nums {
      need := target - n
      if idx, exist := needSet[need]; exist {
        return []int{idx,i}
      } 
      needSet[n] = i
    } 
    return nil
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9)) // → [0,1]
	fmt.Println(twoSum([]int{3,2,4}, 6)) // → [1,2]
	fmt.Println(twoSum([]int{3,3}, 6)) // → [0,1]
}
