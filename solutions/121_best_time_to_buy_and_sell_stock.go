package main

import "fmt"

// #121 Best Time to Buy and Sell Stock [Easy]
// Tags: Array, Dynamic Programming

func maxProfit(prices []int) int {
 // (前-後項)差最多
 // 紀錄當前 max
 // 遇到比較小的就替換基準值
    base := prices[0]
    maxP := 0
    for i := 1; i < len(prices); i++{
      p := prices[i] - base
      if maxP < p {
        maxP = p
      }
      
      if p < 0 {
        base = prices[i]
      }
    }
    return maxP
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4})) // → 5
	fmt.Println(maxProfit([]int{7,6,4,3,1})) // → 0
	fmt.Println(maxProfit([]int{2,4,1})) // → 2
}
