package main

import (
	listnode2 "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

func mergeTwoLists(l1 *listnode2.ListNode, l2 *listnode2.ListNode) *listnode2.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dump := new(listnode2.ListNode)
	if l1.Val < l2.Val {
		dump.Next = l1
		dump.Next.Next = mergeTwoLists(l1.Next, l2)
		return dump.Next
	} else {
		dump.Next = l2
		dump.Next.Next = mergeTwoLists(l1, l2.Next)
		return dump.Next
	}
}
