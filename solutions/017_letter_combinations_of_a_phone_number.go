package main

import "fmt"

// #17 Letter Combinations of a Phone Number [Medium]
// Tags: Hash Table, String, Backtracking

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitToChars := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	result := []string{""}

	for _, digit := range digits {
		temp := []string{}
		chars := digitToChars[string(digit)]

		for _, prevCombination := range result {
			for _, ch := range chars {
				temp = append(temp, prevCombination+string(ch))
			}
		}

		result = temp
	}

	return result
}

func main() {
	fmt.Println(letterCombinations("")) // → []
	fmt.Println(letterCombinations("2")) // → ["a","b","c"]
}
