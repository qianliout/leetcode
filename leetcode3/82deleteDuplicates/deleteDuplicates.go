package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	node := deleteDuplicates(head)
	PrintListNode(node)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates(head.Next)
		return head
	}

	cur := head
	for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
		cur = cur.Next
	}
	return deleteDuplicates(cur.Next)
}
