package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataTree = []struct {
	inputSlice []int
	inputSum   int
	result     bool
}{
	{[]int{}, 0, false},
	{[]int{1}, 1, true},
	{[]int{4, 3, 10}, 7, true},
	{[]int{4, 3, 10}, 10, false},
	{[]int{7, -3, -10, 9, 8, 16}, 4, false},
	{[]int{7, -3, -10, 9, 8, 16}, -6, true},
	{[]int{7, -3, -10, 9, 8, 16}, 24, true},
	{[]int{7, -3, -10, 9, 8, 16}, 32, true},
}

func TestTreePreorder(t *testing.T) {
	for i, data := range testDataTree {
		var root *TreeNode = nil
		for _, number := range data.inputSlice {
			root = root.add(number)
		}
		result := hasPathSum(root, data.inputSum)
		assert.Equal(t, data.result, result, fmt.Sprintf("Test case %d failed: result=%v but must be %v", i, result, data.result))
	}
}
