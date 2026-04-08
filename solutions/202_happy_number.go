package main

import "fmt"

// #202 Happy Number [Easy]
// Tags: Hash Table, Math, Two Pointers

func isHappy(n int) bool {
    
    for true {
      n = dddo(n)
      if n == 1 {
        return true
      }

      if n < 10 {
        return false
      }
    }
    return false
}


func dddo(n int) int {
  total := 0
  
  for true {
    split := n % 10
    total += split * split

    if n / 10 > 0 {
      n /= 10
    } else {
      break
    }
    
  }
  
  return total
}

func main() {
	fmt.Println(isHappy(19)) // → true
	fmt.Println(isHappy(2)) // → false
	fmt.Println(isHappy(1)) // → true
}
