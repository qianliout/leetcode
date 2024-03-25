package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fir, sec, thi := head, head.Next, head.Next.Next

	sec.Next = fir
	fir.Next = swapPairs(thi)
	return sec
}
