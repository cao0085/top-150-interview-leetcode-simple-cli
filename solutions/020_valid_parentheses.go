package main

import "fmt"

// #20 Valid Parentheses [Easy]
// Tags: String, Stack

func isValid(s string) bool {
    dicMap := make(map[string]string,{"(":")","[":"]","{":"}"}

    for i := 1, v := range s {
      before := s[i]

      if v != before

     }
    return false
}

func main() {
	fmt.Println(isValid("()")) // → true
	fmt.Println(isValid("()[]{}")) // → true
	fmt.Println(isValid("(]")) // → false
	fmt.Println(isValid("([])")) // → true
	fmt.Println(isValid("{[]}")) // → true
}
