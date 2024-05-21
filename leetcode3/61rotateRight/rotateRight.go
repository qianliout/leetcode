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
	node := rotateRight(list, 1111)
	PrintListNode(node)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}
	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	k = k % cnt
	if k == 0 {
		return head
	}
	left := cnt - k
	dump := &ListNode{Next: head}
	pre := dump

	for left > 0 {
		pre = pre.Next
		left--
	}

	fist := pre.Next
	pre.Next = nil

	cur = fist
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return fist
}
