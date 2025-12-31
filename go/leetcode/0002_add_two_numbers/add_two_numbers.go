package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func push(listNode *ListNode, value int) *ListNode {
	var nextListNode *ListNode = new(ListNode)
	nextListNode.Val = value
	nextListNode.Next = listNode
	return nextListNode
}

func reverse(listNode *ListNode) *ListNode {
	var reverseList *ListNode
	for listNode != nil {
		reverseList = push(reverseList, listNode.Val)
		listNode = listNode.Next
	}
	return reverseList
}

// 0 ms | Beats 100.00%
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultList *ListNode = nil
	var sum int = 0
	var shift int = 0
	for l1 != nil || l2 != nil {
		var firstVal int
		var secondVal int
		if l1 != nil {
			firstVal = l1.Val
		}
		if l2 != nil {
			secondVal = l2.Val
		}
		sum = firstVal + secondVal + shift
		shift = sum / 10
		sum %= 10
		resultList = push(resultList, sum)
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if shift != 0 {
		resultList = push(resultList, shift)
	}
	return reverse(resultList)
}
