package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataTree = []struct {
	inputSlice   []int
	inputSum     int
	resultSlices [][]int
}{
	{[]int{}, 0, [][]int{}},
	{[]int{1}, 1, [][]int{{1}}},
	{[]int{4, 3, 10}, 7, [][]int{{4, 3}}},
	{[]int{4, 3, 10}, 10, [][]int{}},
	{[]int{7, 4, 9, 3, 5}, 16, [][]int{{7, 4, 5}, {7, 9}}},
}

func TestTreePreorder(t *testing.T) {
	for i, data := range testDataTree {
		var root *TreeNode = nil
		for _, number := range data.inputSlice {
			root = root.add(number)
		}
		result := pathSum(root, data.inputSum)
		assert.Equal(t, data.resultSlices, result, fmt.Sprintf("Test case %d failed: result=%v but must be %v", i, result, data.resultSlices))
	}
}
