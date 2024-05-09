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
	list.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	node := reverseBetween(list, 2, 3)
	PrintListNode(node)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || right <= left || left < 1 {
		return head
	}

	if left == 1 {
		return reverse(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverse(head *ListNode, right int) *ListNode {
	if head == nil || head.Next == nil || right <= 0 {
		return head
	}
	nex := reverse(head.Next, right-1)
	head.Next.Next = head
	head.Next = nil
	return nex
}
