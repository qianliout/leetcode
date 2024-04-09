package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		su := numbers[l] + numbers[r]
		if su == target {
			return []int{l + 1, r + 1}
		} else if su > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}
