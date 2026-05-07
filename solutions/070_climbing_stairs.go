package main

import "fmt"

// #70 Climbing Stairs [Easy]
// Tags: Math, Dynamic Programming, Memoization

func climbStairs(n int) int {
    // 10 可以被幾種只用1,2的排列組合完成?
    // 且不同順序排列都算一種

    if n <= 0 {
      return 0
    } else if n == 1 {
      return 1
    } else if n == 2 {
      return 2
    } else {
      return climbStairs(n-1) + climbStairs(n-2)
    }
}

func main() {
	fmt.Println(climbStairs(1)) // → 1
	fmt.Println(climbStairs(2)) // → 2
	fmt.Println(climbStairs(3)) // → 3
	fmt.Println(climbStairs(10)) // → 89
}
