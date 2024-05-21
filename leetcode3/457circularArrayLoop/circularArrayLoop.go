package main

import (
	"fmt"
)

func main() {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	fmt.Println(circularArrayLoop([]int{-1, 2}))
	fmt.Println(circularArrayLoop([]int{-2, 1, -1, -2, -2}))
	fmt.Println(circularArrayLoop([]int{-2, -3, -9}))
}

func circularArrayLoop1(nums []int) bool {
	for i := range nums {
		if check(nums, i) {
			return true
		}
	}
	return false
}

func check(nums []int, start int) bool {
	cur := start
	flag := nums[start] > 0 // 同向环才能是正确的，中途不能改变方向
	k := 1
	for {
		next := step(cur, nums[cur], len(nums))

		if flag && nums[next] < 0 {
			return false
		}
		if !flag && nums[next] > 0 {
			return false
		}
		if next == start {
			return k > 1
		}
		// 如果检查过程中扫描的数量 kkk 超过了数组长度 nnn，那么根据「鸽笼原理」，必然有数被重复处理了，同时条件一并不符合，因此再处理下去，也不会到达与起点相同的下标，返回 False；
		if k > len(nums) {
			return false
		}
		cur = next
		k++
	}
}

// 使用快慢批针的方式
/*
具体流程是：

在每次移动中，快指针需要走 2 次，而慢指针需要走 1 次；
每次移动的步数等于数组中每个位置存储的元素；
当快慢指针相遇的时候，说明有环。
起始时，让快指针先比慢指针多走一步，当两者在满足题目的两个限制条件的情况下，快满指针能够相遇，则说明有环。

这个题难点在于题目的两个限制条件：

在每次循环的过程中，必须保证所经历过的所有数字都是同号的。
那么，在快指针经历过的每个位置都要判断一下和出发点的数字是不是相同的符号。
当快慢指针相遇的时候，还要判断环的大小不是 1。
所以，找到相遇点的位置后，如果再走 1 步，判断是不是自己。
*/

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		slow := i
		fast := step(slow, nums[slow], len(nums))
		for {
			// 要求要同向，检测快指针两次是否同向就行
			if nums[fast]*nums[i] > 0 && nums[i]*nums[step(fast, nums[fast], len(nums))] > 0 {
				break
			}
			// 慢指针直接走
			slow = step(slow, nums[slow], len(nums))
			// 快指针走两步
			fast = step(fast, nums[fast], len(nums))
			fast = step(fast, nums[fast], len(nums))

			if slow == fast {
				if slow != step(slow, nums[slow], len(nums)) {
					return true
				}
				break
			}
		}
	}
	return false
}

// 这个一定要记住
func step(start, k, n int) int {
	if k > 0 {
		return (start + k) % n
	}
	if k == 0 {
		return start
	}
	if k < 0 {
		// 这里k%n 是容易出错的地方
		return (start + (n + k%n)) % n
	}
	return 0
}
