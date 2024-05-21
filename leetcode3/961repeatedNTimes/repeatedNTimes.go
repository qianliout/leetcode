package main

import (
	"fmt"
)

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}

func repeatedNTimes(nums []int) int {
	ex := make(map[int]int)
	for _, ch := range nums {
		if ex[ch] > 0 {
			return ch
		}
		ex[ch]++
	}
	return 0
}
