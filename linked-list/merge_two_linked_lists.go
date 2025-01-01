package main

import "math"

func main() {

}

func getVal(head *ListNode) int {
	if head == nil {
		return math.MaxInt
	}
	return head.Val
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for list1 != nil || list2 != nil {
		if getVal(list1) < getVal(list2) {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	return dummy.Next
}
