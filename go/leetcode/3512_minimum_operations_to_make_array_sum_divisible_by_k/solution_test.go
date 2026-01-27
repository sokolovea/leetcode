package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	inputSlice []int
	k          int
	result     int
}{
	{[]int{3, 9, 7}, 5, 4},
	{[]int{4, 3, 1}, 4, 0},
	{[]int{3, 2}, 6, 5},
}

func TestSliceSumDivisionMinOps(t *testing.T) {
	for i, data := range testData {
		result := minOperations(data.inputSlice, data.k)
		assert.Equal(t, data.result, result, fmt.Sprintf("Test case %d failed: result for \"%d\" = %d but really is %d!", i, data.inputSlice, result, data.result))
	}
}
