package main

import "fmt"

// #3 Longest Substring Without Repeating Characters [Medium]
// Tags: Hash Table, String, Sliding Window

func lengthOfLongestSubstring(s string) int {
  if(len(s)) == 0 {
      return 0
  }
  left := 0
  max := 0
  currSet := make(map[byte]int)
  for i := 0; i < len(s); i++ {
    c := s[i]
    if oldIndex, exist := currSet[c]; exist && oldIndex >= left {
      left = oldIndex + 1
    }
   currSet[c] = i
    if (i - left) > max {
      max = i - left
    }
  }
  return max + 1
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // → 3
	fmt.Println(lengthOfLongestSubstring("bbbbb")) // → 1
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // → 3
	fmt.Println(lengthOfLongestSubstring("")) // → 0
}
