package main

import (
	"fmt"
)

type taskFunc func([]int, int) []int

// 22 ms | Beats 20.57%
func twoSumSimple(nums []int, target int) []int {
	var sumIndexes []int = []int{0, 0}
out:
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sumIndexes[0] = i
				sumIndexes[1] = j
				break out
			}
		}
	}
	return sumIndexes
}

// 24 ms | Beats 15.38%
func twoSumMap(nums []int, target int) []int {
	var sumIndexes []int = []int{0, 0}
	var numbersMap map[int]int = make(map[int]int) // numbers -> count
	for _, number := range nums {
		numbersMap[number] += 1
	}

	var otherNumberInvariant int
	var minimumRequiredCount int
out:
	for i, number := range nums {
		otherNumberInvariant = target - number
		if otherNumberInvariant == number {
			minimumRequiredCount = 2
		} else {
			minimumRequiredCount = 1
		}
		if numbersMap[number] >= minimumRequiredCount {
			sumIndexes[0] = i
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == otherNumberInvariant {
					sumIndexes[1] = j
					break out
				}
			}
		}
	}
	return sumIndexes
}

// 8 ms | Beats 29.56%
func twoSumMapMap(nums []int, target int) []int {
	var sumIndexes []int = []int{0, 0}
	var numbersMap map[int]map[int]bool = make(map[int]map[int]bool) // numbers -> index

	for index, value := range nums {
		if numbersMap[value] == nil {
			numbersMap[value] = make(map[int]bool)
		}
		numbersMap[value][index] = true
	}
	var currentNumberIndex int = 0
out:
	for number, indexMap := range numbersMap {
		var otherNumberInvariant int = target - number
		var countIndexesInvariant int = len(numbersMap[otherNumberInvariant])
		if ((otherNumberInvariant != number) && countIndexesInvariant > 0) || ((otherNumberInvariant == number) && countIndexesInvariant > 1) {
			for index := range indexMap {
				if currentNumberIndex > 2 {
					break out
				}
				sumIndexes[currentNumberIndex] = index
				currentNumberIndex++
			}
		}
	}
	return sumIndexes
}

// 0 ms | Beats 100.00%
func twoSumOptimized(nums []int, target int) []int {
	var sumIndexes []int = []int{0, 0}
	var numbersMap map[int]int = make(map[int]int) // numbers -> lastIndex

	for index, number := range nums {
		numbersMap[number] = index
	}

	var invariantNumber int
	var invariantIndex int
	for index, number := range nums {
		invariantNumber = target - number
		invariantIndex = numbersMap[invariantNumber]
		// fmt.Printf("Number = %d, invariantNumber = %d\n", number, invariantNumber)
		if invariantIndex != 0 && index != invariantIndex {
			sumIndexes[0] = index
			sumIndexes[1] = invariantIndex
			break
		}
	}
	return sumIndexes
}

// Input: count numbers, then target, then numbers
func main() {

	var taskFunc taskFunc = twoSumOptimized

	var count int
	fmt.Scanf("%d", &count)
	var target int
	fmt.Scanf("%d", &target)
	var numbers []int = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &numbers[i])
	}
	resultArray := taskFunc(numbers, target)
	fmt.Println(resultArray)
}
