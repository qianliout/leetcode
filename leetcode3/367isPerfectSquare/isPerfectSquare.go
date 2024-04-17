package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(4))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(100))
}

func isPerfectSquare(num int) bool {
	if num <= 1 {
		return true
	}

	l, r := 1, num
	for l < r {
		mid := l + (r-l)/2
		if num > mid*mid {
			l = mid + 1
		} else if num < mid*mid {
			r = mid - 1
		} else {
			return true
		}
	}
	return l*l == num
}
