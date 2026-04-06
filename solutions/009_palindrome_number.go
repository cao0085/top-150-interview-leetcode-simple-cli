package main

import "fmt"

// #9 Palindrome Number [Easy]ㄗ/ Tags: Math

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    if x < 10 {
        return true
    }

    result := []int{}
    n := x
    for n > 0 {
        s := n % 10
        result = append([]int{s},result...)
        n /= 10
    }

    for i := 0; i <= len(result)/2; i ++ {
        if result[i] != result[len(result)-1-i] {
            return false
        }
    }
    return true
}

func main() {
	fmt.Println(isPalindrome(121)) // → true
	fmt.Println(isPalindrome(-121)) // → false
	fmt.Println(isPalindrome(10)) // → false
	fmt.Println(isPalindrome(0)) // → true
}
