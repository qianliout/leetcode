package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

// 如果用于排序的话会超时
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = nil
	nex := insertionSortList(next)
	dump := &ListNode{Next: nex}
	pre, cur := dump, dump.Next
	for cur != nil {
		if cur.Val >= head.Val {
			break
		}
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = head
	head.Next = cur

	return dump.Next
}

func sortList(head *ListNode) *ListNode {
	return insertionSortList(head)
}
