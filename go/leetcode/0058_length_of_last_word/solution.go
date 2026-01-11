package main

import (
	"fmt"
	"strings"
)

// 0 ms | Beats 100.00%
func lengthOfLastWord(s string) int {
	var words []string = strings.Split(s, " ")
	var result int = 0
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) != 0 {
			result = len(words[i])
			break
		}
	}
	return result
}

func main() {
	fmt.Printf("%d\n", lengthOfLastWord("Hello from Russia!"))          //7
	fmt.Printf("%d\n", lengthOfLastWord("Russia!"))                     //7
	fmt.Printf("%d\n", lengthOfLastWord(""))                            //0
	fmt.Printf("%d\n", lengthOfLastWord("dfg"))                         //3
	fmt.Printf("%d\n", lengthOfLastWord("   fly me   to   the moon  ")) //4
}
