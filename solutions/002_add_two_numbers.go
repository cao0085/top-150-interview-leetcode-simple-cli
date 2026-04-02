package main

import "fmt"

// #2 Add Two Numbers [Medium]
// Tags: Linked List, Math, Recursion

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  
  difference := len(l1) - len(l2)
  if difference > 0 {
    for range difference {
      l2 = append(l2,0)
     }
  } else if difference < 0 {
    for range (-difference) {
      l1 = append(l1,0)
    }
  }
  anwser := int[]{}
  hold := 0
  
  for i := 0; i < len(l1); i ++ {
    val := l1[i] + l2[i] + hold
    if val >= 10 {
      val = val - 10
      hold = 1
    }
    
    anwser = append(anwser,val)
  }
  // 題目是 linked List:wq
  return anwser
}

func main() {
	// Add test calls here
}
