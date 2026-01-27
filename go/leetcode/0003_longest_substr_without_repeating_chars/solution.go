package main

import "fmt"

// 4 ms | Beats ~70%
func lengthOfLongestSubstring(s string) int {
	var maxLength int = 0
	var currentLength int = 0
	var lettersMap map[rune]bool = make(map[rune]bool)
	var firstGoodIndex int = 0
	for index, letter := range s {
		if lettersMap[letter] {
			maxLength = max(maxLength, currentLength)
			for lettersMap[letter] {
				delete(lettersMap, rune(s[firstGoodIndex]))
				firstGoodIndex += 1
			}
			currentLength = index - firstGoodIndex + 1
		} else {
			currentLength += 1
		}
		lettersMap[letter] = true
	}
	return max(maxLength, currentLength)
}

func main() {
	var inputStr string = " "
	var result int = lengthOfLongestSubstring(inputStr)
	fmt.Printf("%d\n", result)
}
