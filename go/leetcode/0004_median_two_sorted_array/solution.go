package main

import "fmt"

// Speed: O(N); memory: O(N)
// 0 ms | Beats 100.00%
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var numsResult []int = make([]int, len(nums1)+len(nums2))
	var nums1Index int = 0
	var nums2Index int = 0
	var currentNumber int

	for i := range numsResult {
		if nums2Index >= len(nums2) || (nums1Index < len(nums1) && nums1[nums1Index] <= nums2[nums2Index]) {
			currentNumber = nums1[nums1Index]
			nums1Index++
		} else { // nums2Index < len(nums2) invariant
			currentNumber = nums2[nums2Index]
			nums2Index++
		}
		numsResult[i] = currentNumber
	}

	var startIndex int = 0
	var endIndex int = len(numsResult) - 1
	var centerIndexSmaller int = (endIndex - startIndex) / 2
	var centerIndexBigger int = (endIndex - startIndex + 1) / 2

	var resultMedian float64
	if centerIndexSmaller == centerIndexBigger {
		resultMedian = float64(numsResult[centerIndexSmaller])
	} else {
		resultMedian = float64(numsResult[centerIndexSmaller]+numsResult[centerIndexBigger]) / 2.0
	}
	return resultMedian
}

// with bugs: only 1847 / 2098 testcases passed
func findMedianSortedArraysOptimized(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(findMedianAndNeighbours(nums1), findMedianAndNeighbours(nums2))
}

func findMedianAndNeighbours(nums []int) []int {
	var medianSlice []int

	var startIndex int = 0
	var endIndex int = len(nums) - 1

	if len(nums)%2 == 0 {
		medianSlice = make([]int, 2)
		var centerIndexSmaller int = (endIndex - startIndex) / 2
		var centerIndexBigger int = (endIndex - startIndex + 1) / 2
		medianSlice[0] = nums[centerIndexSmaller]
		medianSlice[1] = nums[centerIndexBigger]
	} else {
		if len(nums) == 1 {
			medianSlice = make([]int, 1) //test array without slice ???
			medianSlice[0] = nums[0]
		} else {
			medianSlice = make([]int, 3)
			var centerIndex int = (endIndex - startIndex) / 2
			medianSlice[0] = nums[centerIndex-1]
			medianSlice[1] = nums[centerIndex]
			medianSlice[2] = nums[centerIndex+1]
		}
	}
	return medianSlice
}

func main() {
	var nums1 []int = []int{1, 3}
	var nums2 []int = []int{2}
	var result = findMedianSortedArrays(nums1, nums2)
	fmt.Printf("%v\n", result)
}
