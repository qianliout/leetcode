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

func isPerfectSquare1(num int) bool {
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

func isPerfectSquare2(num int) bool {
	if num <= 1 {
		return true
	}
	left, right := 1, num
	for left < right {
		mid := left + (right-left)>>1
		// 左边界
		if mid < num && mid >= 0 && mid*mid >= num {
			right = mid
		} else {
			left = left + 1
		}
	}
	return left*left == num
}

func isPerfectSquare(num int) bool {
	if num <= 1 {
		return true
	}
	left, right := 1, num
	for left < right {
		mid := left + (right-left+1)>>1
		// 右边界
		if mid < num && mid >= 0 && mid*mid <= num {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left*left == num
}
