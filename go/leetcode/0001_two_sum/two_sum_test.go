package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var data []int = []int{-9, 8, -4, 10, 8, 1}
	var target int = -3
	var result = []int{2, 5}
	got := twoSumSimple(data, target)
	if !slices.Equal(got, result) {
		t.Errorf("TestTwoSum = %v; want %v\n", got, result)
	}
}
