package main

import "fmt"

// #3 Longest Substring Without Repeating Characters [Medium]
// Tags: Hash Table, String, Sliding Window

func lengthOfLongestSubstring(s string) int {
  si := append([]int{0},0)
  currSet := make(map[byte]int)
  for i := 0; i < len(s); i++ {
    c := s[i]
    if oldIndex, exist := recordSet[c]; !exist {
      recordSet[c] = i     
    } else {
      
    }
  }
  return sliceIndex[1] - sliceIndex[0]
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // → 3
	fmt.Println(lengthOfLongestSubstring("bbbbb")) // → 1
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // → 3
	fmt.Println(lengthOfLongestSubstring("")) // → 0
}
