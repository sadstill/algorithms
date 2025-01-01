package main

import (
	"fmt"
)

func main() {
	ln6 := ListNode{9, nil}
	ln5 := ListNode{8, &ln6}
	ln4 := ListNode{6, &ln5}
	ln3 := ListNode{2, &ln4}
	ln2 := ListNode{3, &ln3}
	ln1 := ListNode{4, &ln2}

	ll := middleNodeTask(&ln1)

	for ll != nil {
		fmt.Println(ll)
		ll = ll.Next
	}
}

func middleNodeTask(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
