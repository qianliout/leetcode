package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

// 为啥不得行呢，有问题
func detectCycle(head *ListNode) *ListNode {
	if !hasCycle(head) {
		return nil
	}

	slow, fast := head, head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
