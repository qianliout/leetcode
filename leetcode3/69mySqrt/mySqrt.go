package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/listnode"
)

func main() {
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(16))
}

// 找一个数 mid*mid<=x,找到的最大值，相当于右端点
func mySqrt(x int) int {
	l, r := 0, x+1
	for l < r {
		mid := l + (r-l+1)>>1
		if mid < x+1 && mid > 0 && mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}

func kthToLast(head *ListNode, k int) int {
	all := 0
	cur := head
	for cur != nil {
		all++
		cur = cur.Next
	}
	if k > all {
		k = k % all

	}
	sub := all - k
	cur = head
	for sub > 0 && cur != nil {
		sub--
		cur = cur.Next
	}
	if cur == nil {
		return 0
	}
	return cur.Val
}
