package main

import "fmt"

// #122 Best Time to Buy and Sell Stock II [Medium]
// Tags: Array, Dynamic Programming, Greedy

func maxProfit(prices []int) int {

 // 未持有:看隔天是否漲再持有
 // 已持有:每步算獲利,累積到下個低價再空手
    curr := -1
    profit := 0
    total := 0

    for i := 0; i < len(prices) - 1; i++ {
      
      if curr == -1 {
        if prices[i] < prices[i+1] {
          curr = prices[i]
        }
      } else {
        p := prices[i] - curr
        
        if p > profit {
          profit = p
        }

        if prices[i] > prices[i+1] {
          total += profit
          profit = 0
          curr = -1
        }
        
      }
    }
    
    // 還須補是否最大峰值
    if curr != -1 && prices[len(prices)-1] - curr > 0 {
      total += (prices[len(prices)-1] - curr)
    }
    return total
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4})) // → 7
	fmt.Println(maxProfit([]int{1,2,3,4,5})) // → 4
	fmt.Println(maxProfit([]int{7,6,4,3,1})) // → 0
}
