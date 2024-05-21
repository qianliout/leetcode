package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left >= right {
		return head
	}

	if left == 1 {
		// 找到right
		first := head
		rightNode := head
		for rightNode != nil && right > 0 {
			rightNode = rightNode.Next
			right--
		}
		node := reverse(first, rightNode)
		cur := node
		for cur != nil && cur.Next != nil && cur != rightNode && cur.Next != rightNode {
			cur = cur.Next
		}
		cur.Next = rightNode
		return node

	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverse(head *ListNode, end *ListNode) *ListNode {
	if head == end || head.Next == end {
		return head
	}
	next := reverse(head.Next, end)
	head.Next.Next = head
	head.Next = nil
	return next
}
