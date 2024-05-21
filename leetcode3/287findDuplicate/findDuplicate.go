package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate1(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = 0
	for fast != slow {
		fast, slow = nums[fast], nums[slow]
	}
	return fast
}
