package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//var (
//	head  = node1
//	node1 = &ListNode{1, node2}
//	node2 = &ListNode{2, node3}
//	node3 = &ListNode{3, node4}
//	node4 = &ListNode{4, node5}
//	node5 = &ListNode{5, nil}
//)

//func main() {
//	reorderList(head)
//	for current := head; current != nil; current = current.Next {
//		fmt.Println(current)
//	}
//}

// o(1) решение по памяти
func reorderList(head *ListNode) {
	secondHalf := reverseLinkedList(FindMiddleNode(head))

	l := head
	r := secondHalf
	prev := &ListNode{}
	for r != nil && r != l {
		temp := l.Next
		prev.Next = l
		l.Next = r
		prev = r
		l = temp
		r = r.Next
	}
}

func FindMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
