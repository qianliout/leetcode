package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits(5))
	fmt.Println(countBits(2))
}

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = cnt(i)
	}
	return ans
}

func cnt(n int) int {
	ans := 0
	for n > 0 {
		if n&1 == 1 {
			ans++
		}
		n = n >> 1
	}
	return ans
}

func find2(n int) []int {
	// dp[n]表示n 有多少个1
	dp := make(map[int]int)
	ans := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 0; i <= n; i++ {
		dp[i] = dp[i>>1] + 1&i
		ans[i] = dp[i]
	}

	return ans
}

func isPowerOfFour(n int) bool {
	return n != 0 && n&(n-1) == 0
}
