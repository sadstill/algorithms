package main

func isPalindrome(head *ListNode) bool {
	middleNode := findMiddleNode(head)

	startNodeOfReversedPart := reverseNode(middleNode)

	p1 := head
	p2 := startNodeOfReversedPart
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	reverseNode(startNodeOfReversedPart)

	return true
}

func findMiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseNode(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
