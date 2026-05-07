package main

import "fmt"

// #134 Gas Station [Medium]
// Tags: Array, Greedy

func canCompleteCircuit(gas []int, cost []int) int {

  // 條件A: 當 gas >= cost 才可以環繞一圈
  // 條件B: 假設有一個 currGas[] 紀錄每一步的貢獻值
  //  - 第一步必定選正貢獻值 currGas[正]
  //  - 直到"跑完一圈"或發生"Sum(currGas[正,?,?,?,?]) < 0"
  //  - Sum < 0 代表目前走過的路都不能當作起始點，原因
  //    1. 首次發生 Sum(currGas) + Next < 0 時, Next 貢獻值必為負
  //    2. 任一挑選一個點開始 = 任意分兩段(currGas[:?],currGas[?:])
  //    3. 已知 Sum(currGas[:?]) 必為正數 = 捨棄部分為正數 = 剩餘加總不可能大於 Next 貢獻值

  total := 0;
  currG := 0;
  start := 0;
  
  for i := 0; i < len(gas); i ++ {
    diff := gas[i] - cost[i]
    total += diff
    currG += diff

    if currG < 0 {
      start = i + 1
      currG = 0
    }
  }

  if total < 0 {
    return -1
  }

  if currG > 0 {
    return start
  }
  return -1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2})) // → 3
	fmt.Println(canCompleteCircuit([]int{2,3,4}, []int{3,4,3})) // → -1
}
