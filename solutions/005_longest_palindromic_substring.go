package main

import "fmt"

// #5 Longest Palindromic Substring [Medium]
// Tags: String, Dynamic Programming

func longestPalindrome(s string) string {
    
    if len(s) == 1 {
      return s
    } else if len(s) == 0 {
      return ""
    }

    newS := []int{len(s)/2,len(s)/2}
    
    for i := range s {
      ss := s[0:i+1]
      n := len(ss)
      c := n / 2
      left, right := c, c

      if n % 2 == 0 {
        left = c - 1
      }

      for left >= 0 && right < n && ss[left] == ss[right] {
        if (newS[1] - newS[0]) < (right - left) {
          newS = []int{left, right}
        }
        left--
        right++
     }
   }

   return s[newS[0]:newS[1]+1]
}

func main() {
	fmt.Println(longestPalindrome("babad")) // → "bab"
	fmt.Println(longestPalindrome("cbbd")) // → "bb"
	fmt.Println(longestPalindrome("a")) // → "a"
}
