package main

import "fmt"

// #21 Merge Two Sorted Lists [Easy]
// Tags: Linked List, Recursion

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    dummy := &ListNode{Val: 0, Next: nil}
    curr := dummy
    
    for list1 != nil && list2 != nil {
      if list1.Val >= list2.Val {
        curr.Next = list2
        list2 = list2.Next
      } else {
        curr.Next = list1
        list1 = list1.Next
      }

      curr = curr.Next
    }

      for node := dummy.Next; node != nil; node = node.Next {
          fmt.Print("sssss",node.Val, " ")
      }
    return nil
}

func main() {
	// Add test calls here
      // ✅ 創建 list1: 1 -> 3
      list1 := &ListNode{Val: 1}
      list1.Next = &ListNode{Val: 3}

      // ✅ 創建 list2: 2 -> 4
      list2 := &ListNode{Val: 2}
      list2.Next = &ListNode{Val: 4}

      // 合併
      result := mergeTwoLists(list1, list2)
      fmt.Println("Result:", result)
}
