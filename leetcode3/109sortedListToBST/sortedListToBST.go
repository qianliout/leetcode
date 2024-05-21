package main

import (
	. "outback/geeke/leetcode/common/listnode"
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	dump := &ListNode{Next: head}
	pre, slow, fast := dump, head, head
	for fast != nil && fast.Next != nil {
		pre = pre.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{Val: slow.Val}
	pre.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}
