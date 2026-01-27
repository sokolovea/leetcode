package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataShiftSlice = []struct {
	inputSlice      []int
	startShiftIndex int
	resultSlice     []int
}{
	{[]int{}, 0, []int{}},
	{[]int{}, 1, []int{}},
	{[]int{}, -1, []int{}},
	{[]int{}, 2, []int{}},
	{[]int{1}, 0, []int{1}},
	{[]int{1}, 1, []int{1}},
	{[]int{1}, 50, []int{1}},
	{[]int{1}, -50, []int{1}},
	{[]int{1, 2}, 0, []int{2, 2}},
	{[]int{1, 2}, 1, []int{1, 2}},
	{[]int{1, 2}, -1, []int{1, 2}},
	{[]int{1, 1, 2}, 0, []int{1, 2, 2}},
	{[]int{1, 1, 2}, 1, []int{1, 2, 2}},
	{[]int{1, 1, 2}, 2, []int{1, 1, 2}},
	{[]int{1, 1, 2}, 3, []int{1, 1, 2}},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 0, []int{0, 1, 1, 1, 2, 2, 3, 3, 4, 4}},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 0, 1, 1, 1, 2, 3, 3, 4, 4}},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 8, []int{0, 0, 1, 1, 1, 2, 2, 3, 4, 4}},
}

func TestShift(t *testing.T) {
	for i, data := range testDataShiftSlice {
		processedSlice := data.inputSlice
		t.Logf("Input slice %v", processedSlice)
		startShiftIndex := data.startShiftIndex
		shiftSlice(processedSlice, startShiftIndex)
		expectedSlice := data.resultSlice
		assert.Equal(t, expectedSlice, processedSlice, fmt.Sprintf("Test case %d failed: result slice %v but must be %v", i, processedSlice, expectedSlice))
	}
}

var testDataNums = []struct {
	nums       []int
	k          int
	resultNums []int
}{
	{[]int{1, 2}, 2, []int{1, 2}},                                  //[1,2]
	{[]int{1, 1, 2}, 2, []int{1, 2}},                               //[1,2,_]
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}}, //[0,1,2,3,4,_,_,_,_,_]
}

type RemoveDuplicates func(nums []int) int

func InternalTestRemoveDuplicates(f RemoveDuplicates, t *testing.T) {
	for i, data := range testDataNums {
		inputArray := data.nums
		t.Logf("Input slice %v", inputArray)
		result := f(inputArray)
		expectedK := data.k
		expectedArray := data.resultNums
		assert.Equal(t, expectedK, result, fmt.Sprintf("Test case %d failed: result %v but must be %d", i, result, expectedK))
		assert.Equal(t, expectedArray[:expectedK], inputArray[:expectedK], fmt.Sprintf("Test case %d failed: result slice %v but must be %v", i, inputArray[:expectedK], expectedArray[:expectedK]))
	}
}

func TestRemoveDuplicatesDummy(t *testing.T) {
	InternalTestRemoveDuplicates(removeDuplicatesDummy, t)
}

func TestRemoveDuplicatesOptimized(t *testing.T) {
	InternalTestRemoveDuplicates(removeDuplicatesOptimized, t)
}
