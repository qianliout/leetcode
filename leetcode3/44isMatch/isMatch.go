package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "*"))
}

func isMatch(s string, p string) bool {
	ss, pp := []byte(s), []byte(p)
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(ss); i++ {
		dp[i] = make([]bool, len(pp)+1)
	}
	// 初值
	dp[0][0] = true
	for j := 1; j <= len(pp); j++ {
		if pp[j-1] != '*' {
			break
		}
		dp[0][j] = true
	}

	for i := 1; i <= len(ss); i++ {
		for j := 1; j <= len(pp); j++ {
			if ss[i-1] == pp[j-1] || pp[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if pp[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[len(ss)][len(pp)]
}
