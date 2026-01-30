package main

import (
	"math"
	"sort"
)

func thirdMaxDummy(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var maxNums []int = make([]int, 1, 3)
	maxNums[0] = nums[0]
	for _, num := range nums {
		if len(maxNums) == 3 {
			break
		}
		if num != maxNums[len(maxNums)-1] {
			maxNums = append(maxNums, num)
		}
	}
	if len(maxNums) < 3 {
		return maxNums[0]
	}
	return maxNums[2]
}

func thirdMaxOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var maxNums []int = make([]int, 0, 3)
	for counter := range 3 {
		currentMax := math.MinInt
		isFound := false
		for _, num := range nums {
			if num > currentMax && (counter == 0 || num < maxNums[counter-1]) {
				currentMax = num
				isFound = true
			}
		}
		if isFound {
			maxNums = append(maxNums, currentMax)
		} else {
			break
		}
	}
	if len(maxNums) < 3 {
		return maxNums[0]
	}
	return maxNums[2]
}
