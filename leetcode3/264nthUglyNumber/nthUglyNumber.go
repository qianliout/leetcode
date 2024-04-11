package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(nthUglyNumber(1))
	fmt.Println(nthUglyNumber(2))
	fmt.Println(nthUglyNumber(3))
	fmt.Println(nthUglyNumber(4))
	fmt.Println(nthUglyNumber(5))
	fmt.Println(nthUglyNumber(6))
	fmt.Println(nthUglyNumber(7))
	fmt.Println(nthUglyNumber(8))
	fmt.Println(nthUglyNumber(9))
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(11))
}

func nthUglyNumber(n int) int {
	a2, a3, a5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	exit := make(map[int]bool)
	cnt := 1
	for cnt < n {
		ans := Min(dp[a2]*2, dp[a3]*3, dp[a5]*5)
		if !exit[ans] {
			dp[cnt] = ans
			exit[ans] = true
			cnt++
		}
		if ans == dp[a2]*2 {
			a2++
		}
		if ans == dp[a3]*3 {
			a3++
		}
		if ans == dp[a5]*5 {
			a5++
		}
	}

	return dp[n-1]
}
