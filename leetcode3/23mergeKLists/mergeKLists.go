package main

import (
	. "outback/geeke/leetcode/common/listnode"
)

func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2

	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return merge2(left, right)
}

func merge2(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = merge2(list1.Next, list2)
		return list1
	} else {
		list2.Next = merge2(list1, list2.Next)
		return list2
	}
}
