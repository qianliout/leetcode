package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(math.MinInt32))
	fmt.Println(reverse(math.MaxInt32))
	fmt.Println(3 ^ -2)
	fmt.Println(-3 ^ 2)
	fmt.Println(3 ^ 2)
}

func reverse2(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	num := 0

	for x != 0 {
		y := x % 10
		if num*10+y > math.MaxInt32 {
			return 0
		}

		num = num*10 + y
		x = x / 10
	}
	return num
}

func reverse(x int) int {
	ans := 0
	fl := true
	if x < 0 {
		fl = false
	}
	for x != 0 {
		// 如果 x 是负数，那么 y 也是负数
		y := x % 10
		// 因为不能存储32位以上的，所以这样才是正解
		if (fl && ans > (math.MaxInt32-y)/10) || (!fl && ans < (math.MinInt32-y)/10) {
			return 0
		}

		ans = ans*10 + y
		x = x / 10
	}
	return ans
}
