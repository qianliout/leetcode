package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	node := deleteDuplicates(head)
	PrintListNode(node)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := deleteDuplicates(head.Next)
	if head.Val == next.Val {
		head.Next = next.Next
		return head
	}
	head.Next = next
	return head
}
