package main

func arraySum(nums []int) int {
	var sum int = 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func minOperations(nums []int, k int) int {
	return arraySum(nums) % k
}
