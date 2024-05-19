package main

import (
	"fmt"
)

func main() {
	fmt.Println(validMountainArray([]int{0, 2, 3, 4, 5, 5, 4, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	mx := findMax(arr)
	if mx == 0 || mx == len(arr)-1 {
		return false
	}

	for i := mx; i > 0; i-- {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	for i := mx; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}

func findMax(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return i
		}
	}
	return len(arr) - 1
}
