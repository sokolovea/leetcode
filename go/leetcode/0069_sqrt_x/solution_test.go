package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	number    int
	root      int
	isPerfect bool
}{
	{0, 0, true},
	{1, 1, true},
	{2, 1, false},
	{3, 1, false},
	{4, 2, true},
	{8, 2, false},
	{9, 3, true},
	{99, 9, false},
	{100, 10, true},
	{102, 10, false},
	{119, 10, false},
	{10003, 100, false},
}

func TestMySqrt(t *testing.T) {
	for i, data := range testData {
		result := mySqrt(data.number)
		assert.Equal(t, data.root, result, fmt.Sprintf("Test case %d failed: for %v result %v but should be %v", i, data.number, result, data.root))
	}
}

func TestIsPerfectSquare(t *testing.T) {
	for i, data := range testData {
		result := isPerfectSquare(data.number)
		assert.Equal(t, data.isPerfect, result, fmt.Sprintf("Test case %d failed: for %v result %v but should be %v", i, data.number, result, data.isPerfect))
	}
}
