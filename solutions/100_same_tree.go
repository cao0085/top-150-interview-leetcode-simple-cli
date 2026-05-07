package main

import "fmt"

// #100 Same Tree [Easy]
// Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

    if p == nil && q == nil {
      return true
    }

    if p == nil || q == nil { 
     return false
    }
    
    if p.Val != q.Val { 
     return false
    }
   
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// Add test calls here
}
