package main

import "fmt"

// #14 Longest Common Prefix [Easy]
// Tags: String, Trie

func longestCommonPrefix(strs []string) string {
    
    ss := strs[0]

    for _, s := range strs {
      
     if len(s) < len(ss) {
        ss = s
     }
    }

    i := 0
    for i < len(ss) {
      target := ss[i]
      cc := true
      for _, v := range strs {
        if target != v[i] {
          cc = false
          break;
        }
      }
      
      if cc {
        i ++
      } else {
        break;
      }
    }
    
    return ss[:i]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"})) // → "fl"
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"})) // → ""
	fmt.Println(longestCommonPrefix([]string{"interview"})) // → "interview"
}
