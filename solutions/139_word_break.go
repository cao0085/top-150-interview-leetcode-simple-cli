package main

import "fmt"

// #139 Word Break [Medium]
// Tags: Array, Hash Table, String, Dynamic Programming, Trie, Memoization

func wordBreak(s string, wordDict []string) bool {
    
    dictSet := make(map[string]bool)
    dp := make([]bool,len(s))

    for _, v := range wordDict {
     dictSet[v] = false
    }
    
    for i := 0; i < len(s); i++ {
      
      ss := s[:i+1]

      if _, exsited := dictSet[ss]; exsited {
        dp[i] = true
        continue
      }

      for j, vv := range dp {
        
       if vv == true {
        //fmt.Println(dp,i)
        need := s[j+1:i+1]
        //fmt.Println(need)
        if _, exsited := dictSet[need]; exsited {
          dp[i] = true
        }
       }
      }
    }

    return dp[len(s)-1]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet","code"})) // → true
	fmt.Println(wordBreak("applepenapple", []string{"apple","pen"})) // → true
	fmt.Println(wordBreak("catsandog", []string{"cats","dog","sand","and","cat"})) // → false
}
