package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = &ListNode{Val: 5}
	// list.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	node := reverseKGroup(list, 3)
	PrintListNode(node)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	d := k
	last := head
	for last != nil && d > 0 {
		last = last.Next
		d--
	}
	if d > 0 {
		return head
	}

	node := reverse(head, last)

	nex := reverseKGroup(last, k)

	cur := node
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = nex
	return node
}

func reverse(head, end *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next == end {
		return head
	}
	nex := reverse(head.Next, end)
	head.Next.Next = head
	head.Next = nil
	return nex
}
