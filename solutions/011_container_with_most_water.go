package main

import "fmt"

// #11 Container With Most Water [Medium]
// Tags: Array, Two Pointers, Greedy

func maxArea(height []int) int{

    left := 0
    right := len(height) - 1
    max := 0

    for left <= right {
      
      h := height[left]
      if height[right] < h {
        h = height[right]
      }

      w := right - left
      if max < h * w {
         max = h * w
      }

      if height[left] > height[right] {
        right --
      } else {
        left ++
      }
    }
    return max
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7})) // → 49
	fmt.Println(maxArea([]int{1,1})) // → 1
}
