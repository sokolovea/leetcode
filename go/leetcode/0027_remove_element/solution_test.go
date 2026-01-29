package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	inputSlice  []int
	inputVal    int
	resultSlice []int
	resultK     int
}{
	{[]int{}, 2, []int{}, 0},
	{[]int{1}, 0, []int{1}, 1},
	{[]int{1, 2}, 1, []int{2}, 1},
	{[]int{1, 2, 3, 2, 1, 2, 1}, 1, []int{2, 2, 2, 3}, 4},
	{[]int{1, 2, 3, 2, 1, 2, 1}, 1, []int{3, 2, 2, 2}, 4}, //result slice elements may be in any order
}

func TestRemoveElements(t *testing.T) {
	for i, data := range testData {
		inputSliceCopy := make([]int, len(data.inputSlice))
		copy(inputSliceCopy, data.inputSlice)
		result := removeElement(data.inputSlice, data.inputVal)
		assert.ElementsMatch(t, data.resultSlice, data.inputSlice[0:result], fmt.Sprintf("Test case %d failed: for %v result %v but should be %v", i, data.inputSlice[0:result], result, data.resultSlice))
	}
}
