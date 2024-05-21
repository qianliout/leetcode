package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(integerReplacement(4))
	fmt.Println(integerReplacement(5))
	fmt.Println(integerReplacement(6))
	fmt.Println(integerReplacement(7))
	fmt.Println(integerReplacement(8))
}

// 内存会超
func integerReplacement2(n int) int {
	dp := make([]int, n+10)
	dp[0], dp[1], dp[2], dp[3] = 0, 0, 1, 2

	for i := 4; i <= n; i++ {
		if i&1 == 0 {
			dp[i] = dp[i>>1] + 1
		} else {
			add := dp[(i+1)>>1]
			sub := dp[(i-1)>>1]
			dp[i] = Min(add, sub) + 2
		}
	}

	return dp[n]
}
func integerReplacement(n int) int {
	res := 0
	for n > 1 {
		// 特定
		if n <= 3 {
			res = res + n - 1
			break
		}

		if n%4 == 3 {
			res++
			n = n + 1
		}
		if n%4 == 1 {
			res++
			n = n - 1
		}
		if n%2 == 0 {
			res++
			n = n / 2
		}
	}
	return res
}
