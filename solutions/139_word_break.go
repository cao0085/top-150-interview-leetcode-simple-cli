package main

import "fmt"

// #139 Word Break [Medium]
// Tags: Array, Hash Table, String, Dynamic Programming, Trie, Memoization

func wordBreak(s string, wordDict []string) bool {
    return false
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet","code"})) // → true
	fmt.Println(wordBreak("applepenapple", []string{"apple","pen"})) // → true
	fmt.Println(wordBreak("catsandog", []string{"cats","dog","sand","and","cat"})) // → false
}
