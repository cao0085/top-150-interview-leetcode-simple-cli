package main

import "fmt"

// #71 Simplify Path [Medium]
// Tags: String, Stack

func simplifyPath(path string) string {

    if path[0] != '/' {
      return "error"
    }
    
    stack := []string{}
    currS := ""
    
    for i := range path{
      s := path[i]
      if s == '/' {
        stack = append(stack, currS)
        currS = ""
      } else {
        currS += string(s)
      }
    }

    stack = append(stack, currS)
    
    result := ""

    for i := 0; i < len(stack) - 1; i++{
      if stack[i] == ".." || stack[i+1] == ".." {
        continue
      } else if stack[i] == ""{

      } else {
        result += "/"
        result += stack[i]
      }
    }
    
    if stack[len(stack)-1] != "" {
         result += "/"
        result += stack[len(stack)-1]
    }

    return result
}

func main() {
	fmt.Println(simplifyPath("/home/")) // → "/home"
	fmt.Println(simplifyPath("/home//foo/")) // → "/home/foo"
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures")) // → "/home/user/Pictures"
	fmt.Println(simplifyPath("/../")) // → "/"
}
