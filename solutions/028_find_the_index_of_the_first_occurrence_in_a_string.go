package main

import "fmt"

// #28 Find the Index of the First Occurrence in a String [Easy]
// Tags: String, Two Pointers, String Matching

func strStr(haystack string, needle string) int {
    
    var startI int
    ni := 0

    for i, v := range haystack {
        
      if v == rune(needle[ni]) {
        if ni == 0 {
          startI = i
        }
        ni ++
      } else {
        ni = 0
      }

      if ni == len(needle) {
        return startI
       }
    }
    return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad")) // → 0
	fmt.Println(strStr("leetcode", "leeto")) // → -1
	fmt.Println(strStr("hello", "ll")) // → 2
}
