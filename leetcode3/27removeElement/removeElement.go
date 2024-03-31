package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	start := 0
	for _, nu := range nums {
		if nu != val {
			nums[start] = nu
			start++
		}
	}
	return start
}
