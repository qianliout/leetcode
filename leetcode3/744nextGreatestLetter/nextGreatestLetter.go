package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreatestLetter([]byte("XXYY"), 'Z'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)>>1
		if mid >= 0 && mid < len(letters) && letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < 0 || left >= len(letters) || letters[left] <= target {
		return letters[0]
	}
	return letters[left]
}
