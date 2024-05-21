package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	k := 20
	for k > 0 {
		if a == nil {
			a = headB
			k--
		}
		if b == nil {
			b = headA
			k--
		}
		a = a.Next
		b = b.Next

		if a == b {
			return a
		}
	}

	return nil
}
