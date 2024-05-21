package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	dump := head
	// first1 := dump
	// first2 := dump.Next

	end1 := dump
	end2 := dump.Next
	for end2 != nil && end2.Next != nil {
		end1 = end1.Next
		end2 = end2.Next
	}
	end1.Next = nil
	first := head.Next
	reorderList(first)

	head.Next = end2
	head.Next.Next = first
}

func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	pre, last := head, head.Next
	for last != nil && last.Next != nil {
		pre = pre.Next
		last = last.Next
	}
	last.Next = head.Next
	head.Next = last
	pre.Next = nil
	reorderList2(last.Next)
}
