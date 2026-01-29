package main

func removeElement(nums []int, val int) int {
	index := 0
	lastIndex := len(nums) - 1
	for index <= lastIndex {
		if nums[index] == val {
			nums[index] = nums[lastIndex]
			lastIndex--
		} else {
			index++
		}
	}
	return lastIndex + 1
}
