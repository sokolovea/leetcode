package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataShiftSlice = []struct {
	inputFirst  []int
	inputSecond []int
	result      []int
}{
	{[]int{}, []int{}, []int{}},
	{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
	{[]int{2, 10}, []int{1, 2}, []int{1, 2, 2, 10}},
	{[]int{1, 2}, []int{2, 10}, []int{1, 2, 2, 10}},
	{[]int{1, 4, 9, 12, 19, 23}, []int{-9, 13, 17, 17, 17}, []int{-9, 1, 4, 9, 12, 13, 17, 17, 17, 19, 23}},
}

type mergeTwoListsType func(*ListNode, *ListNode) *ListNode

func ReverseListInnerTest(t *testing.T, mergeFunc mergeTwoListsType) {
	for i, data := range testDataShiftSlice {
		var firstList *ListNode = nil
		var secondList *ListNode = nil
		for _, number := range data.inputFirst {
			tempList := new(ListNode)
			tempList.Next = firstList
			tempList.Val = number
			firstList = tempList
		}
		for _, number := range data.inputSecond {
			tempList := new(ListNode)
			tempList.Next = secondList
			tempList.Val = number
			secondList = tempList
		}

		firstList = reverseList(firstList)
		secondList = reverseList(secondList)

		resultList := mergeFunc(firstList, secondList)

		for index, number := range data.result {
			assert.NotNil(t, resultList, fmt.Sprintf("Test case %d failed: result slice is nil", i))
			listNumber := resultList.Val
			resultList = resultList.Next
			assert.Equal(t, number, listNumber, fmt.Sprintf("Test case %d failed: resultList[%d]=%v but must be %v", i, index, listNumber, number))
		}
	}
}

func TestReverseList(t *testing.T) {
	ReverseListInnerTest(t, mergeTwoLists)
}

func TestReverseListRecursive(t *testing.T) {
	ReverseListInnerTest(t, mergeTwoListsRecursive)
}
