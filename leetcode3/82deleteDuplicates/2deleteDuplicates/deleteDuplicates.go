package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

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
	for cur != nil {
		if cur.Next == nil || cur.Val != cur.Next.Val {
			break
		}
		cur = cur.Next
	}
	return deleteDuplicates(cur.Next)
}
