package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//func main() {
//	ln6 := ListNode{9, nil}
//	ln5 := ListNode{8, &ln6}
//	ln4 := ListNode{6, &ln5}
//	ln3 := ListNode{2, &ln4}
//	ln2 := ListNode{3, &ln3}
//	ln1 := ListNode{4, &ln2}
//
//	ll := removeNthFromEnd(&ln1, 6)
//
//	for ll != nil {
//		fmt.Println(ll)
//		ll = ll.Next
//	}
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	fast := dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	slow := dummy
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
