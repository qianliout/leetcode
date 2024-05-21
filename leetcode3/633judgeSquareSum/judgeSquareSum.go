package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum2(2147482647))
}

func judgeSquareSum2(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		ans := left*left + right*right
		if ans == c {
			return true
		} else if ans < c {
			left++
		} else if ans > c {
			right--
		}
	}

	return false
}

func judgeSquareSum(c int) bool {
	left, right := 1, c/2
	for left <= right {
		ans := left*left + right*right
		if ans == c {
			return true
		} else if ans < c {
			left++
		} else if ans > c {
			right--
		}
	}

	return false
}
