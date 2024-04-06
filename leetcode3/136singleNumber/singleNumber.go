package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{1, 1, 2, 2, 3}))
}

func singleNumber(nums []int) int {
	ans := 0
	for _, ch := range nums {
		ans = ans ^ ch
	}
	return ans
}
