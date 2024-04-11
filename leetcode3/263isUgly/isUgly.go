package main

import (
	"fmt"
)

func main() {
	fmt.Println(isUgly(14))
}

func isUgly2(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 || n == 2 || n == 3 || n == 5 {
		return true
	}
	if n%2 == 0 {
		return isUgly2(n / 2)
	}
	if n%3 == 0 {
		return isUgly2(n / 3)
	}
	if n%5 == 0 {
		return isUgly2(n / 5)
	}

	return false
}
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	ugly := []int{2, 3, 4}

	for _, ch := range ugly {
		for n%ch == 0 {
			n = n / ch
		}
	}

	return n == 1
}
