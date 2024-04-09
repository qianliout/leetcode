package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPrimes(5000000))
	fmt.Println(countPrimes2(5000000))
}

func countPrimes(n int) int {
	dp := make([]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = true // 先假设都是质数
	}
	for i := 2; i*i < n; i++ {
		if dp[i] {
			for j := i * i; j < n; j = j + i {
				dp[j] = false
			}
		}
	}
	ans := 0

	for i := 2; i < n; i++ {
		if dp[i] {
			ans++
		}
	}
	return ans
}

func countPrimes2(n int) int {
	dp := make([]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = true // 先假设都是质数
	}
	for i := 2; i < n; i++ {
		if dp[i] {
			for j := i * 2; j < n; j = j + i {
				dp[j] = false
			}
		}
	}
	ans := 0

	for i := 2; i < n; i++ {
		if dp[i] {
			ans++
		}
	}
	return ans
}
