package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	PrintListNode(oddEvenList(head))

}

func oddEvenList2(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil || head.Next.Next == nil {
	// 	return head
	// }
	first, second := &ListNode{}, &ListNode{}
	fir, sec := first, second

	for head != nil {
		fir.Next = head
		fir = fir.Next
		head = head.Next
		if head == nil {
			break
		}
		sec.Next = head
		sec = sec.Next
		head = head.Next
	}
	sec.Next = nil
	fir.Next = second.Next

	return first.Next
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd, event := head, head.Next
	eventStart := event
	for event != nil && event.Next != nil {
		odd.Next = event.Next
		odd = odd.Next
		event.Next = odd.Next
		event = event.Next
	}
	odd.Next = eventStart
	return odd
}
