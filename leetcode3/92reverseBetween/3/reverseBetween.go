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
	node := reverseBetween(list, 2, 4)
	PrintListNode(node)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || right <= left || left < 1 {
		return head
	}

	if left == 1 {
		rightNode := head
		for rightNode != nil && right > 0 {
			rightNode = rightNode.Next
			right--
		}
		lst := reverse(head, rightNode)

		cur := lst
		for cur != nil && cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = rightNode
		return lst
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
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
