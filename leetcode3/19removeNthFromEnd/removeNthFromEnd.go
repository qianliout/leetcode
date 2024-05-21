package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	node := removeNthFromEnd(list, 1)
	PrintListNode(node)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 计算总数
	cnt := 0
	pre := head
	for pre != nil {
		pre = pre.Next
		cnt++
	}

	sub := cnt - n
	dump := &ListNode{Next: head}
	dump = dfs(dump, sub)
	return dump.Next
}

func bfs(head *ListNode, n int) *ListNode {

	if n == 0 || head == nil {
		return head
	}
	// 计算总数
	cnt := 0
	pre := head
	for pre != nil {
		pre = pre.Next
		cnt++
	}

	sub := cnt - n

	dump := &ListNode{Next: head}
	cur := dump

	for sub > 0 {
		cur = cur.Next
		sub--
	}
	if sub == 0 && cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
	}

	return dump.Next
}

// 删除开始的第n个节点
// 第一个是 dump 节点，从 dump 接口的下一个节点开始算
func dfs(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if n == 0 {
		if head != nil && head.Next != nil {
			head.Next = head.Next.Next
		}
		return head
	}
	head.Next = dfs(head.Next, n-1)
	return head
}
