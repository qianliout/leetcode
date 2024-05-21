package main

import (
	"fmt"
	"math"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}

func nthSuperUglyNumber(n int, primes []int) int {
	idx := make([]int, len(primes))
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		ugly := math.MaxInt64
		for k, pri := range primes {
			ugly = Min(ugly, pri*dp[idx[k]])
		}
		dp[i] = ugly

		for k := range idx {
			if dp[i] == dp[idx[k]]*primes[k] {
				idx[k]++
			}
		}
	}

	return dp[n-1]
}
