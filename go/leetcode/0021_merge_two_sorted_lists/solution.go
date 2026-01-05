package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(inputList *ListNode) *ListNode {
	var resultHead *ListNode = nil
	for inputList != nil {
		tempNode := new(ListNode)
		tempNode.Val = inputList.Val
		tempNode.Next = resultHead
		resultHead = tempNode
		inputList = inputList.Next
	}
	return resultHead
}

// 0 ms | Beats 100.00%
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var resultHead *ListNode = nil
	for list1 != nil || list2 != nil {

		tempNode := new(ListNode)
		var tempVal int
		var isLeft bool = true

		if list1 != nil && list2 != nil {
			tempVal = list1.Val
			isLeft = true
			if list2.Val < tempVal {
				isLeft = false
				tempVal = list2.Val
			}
		} else if list1 != nil {
			tempVal = list1.Val
			isLeft = true
		} else if list2 != nil {
			tempVal = list2.Val
			isLeft = false
		}

		if isLeft {
			list1 = list1.Next
		} else {
			list2 = list2.Next
		}

		tempNode.Val = tempVal
		tempNode.Next = resultHead
		resultHead = tempNode
	}
	return reverseList(resultHead)
}

// 0 ms | Beats 100.00%
func mergeTwoListsRecursiveInner(list1 *ListNode, list2 *ListNode, resultList **ListNode) {

	if (list1 == nil) && (list2 == nil) {
		return
	}

	newNode := new(ListNode)
	newNode.Next = *resultList
	*resultList = newNode

	if list1 == nil {
		(*resultList).Val = list2.Val
		list2 = list2.Next
	} else if list2 == nil {
		(*resultList).Val = list1.Val
		list1 = list1.Next
	} else {
		tempVal := list1.Val
		isLeft := true
		if list2.Val < tempVal {
			isLeft = false
			tempVal = list2.Val
		}

		(*resultList).Val = tempVal

		if isLeft {
			list1 = list1.Next
		} else {
			list2 = list2.Next
		}
	}
	mergeTwoListsRecursiveInner(list1, list2, resultList)
}

func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	var resultList *ListNode = nil
	mergeTwoListsRecursiveInner(list1, list2, &resultList)
	return reverseList(resultList)
}
