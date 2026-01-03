package main

func shiftSlice(nums []int, startIndex int) {
	if startIndex < 0 {
		return
	}
	for i := startIndex; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

// O(N^2)
// 256 ms | Beats 5.40%
func removeDuplicatesDummy(nums []int) int {
	var uniqueNumbersCount int = 0
	if len(nums) <= 1 {
		uniqueNumbersCount = len(nums)
	} else {
		uniqueNumbersCount = 1
		var deltaNumsLength int = 0
		var prevNumber int = nums[0]
		for i := 1; i < len(nums)-deltaNumsLength; i++ {
			currentNumber := nums[i]
			if prevNumber == currentNumber {
				shiftSlice(nums[:len(nums)-deltaNumsLength], i) // optimization with [:] (500 ms -> 250 ms)
				i--
				deltaNumsLength++
			} else {
				prevNumber = nums[i]
				uniqueNumbersCount++
			}
		}
	}
	return uniqueNumbersCount
}

// O(N)
// 0 ms | Beats 100.00%
func removeDuplicatesOptimized(nums []int) int {
	var numsSet map[int]bool = make(map[int]bool)
	var realSliceIndex int = 0
	for keyIndex, number := range nums {
		if !numsSet[number] {
			numsSet[number] = true
			nums[realSliceIndex] = nums[keyIndex]
			realSliceIndex++
		}
	}
	return len(numsSet)
}
