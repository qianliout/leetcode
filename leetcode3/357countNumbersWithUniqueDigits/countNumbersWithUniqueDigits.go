package main

import (
	"fmt"
)

func main() {
	fmt.Println(countNumbersWithUniqueDigits(2))
}

func countNumbersWithUniqueDigits(n int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		num := 9 // 最高位可以填充的数字个数(因为没有前导0)
		a := 9   // 每一位可填充的数字个数(可以填充0)
		for j := i - 1; j > 0; j-- {
			num = num * a
			a--
		}
		dp[i] = dp[i-1] + num
	}
	return dp[n]
}
