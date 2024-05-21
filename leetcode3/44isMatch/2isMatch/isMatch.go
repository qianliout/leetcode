package main

import (
	"fmt"
)

func main() {
	// fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("a", "*a"))
}

/*
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符序列（包括空字符序列）。
*/
func isMatch2(s string, p string) bool {
	s = " " + s
	p = " " + p
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	// 这个特例不能不加，但是很容易忘记，要特别注意
	for j := 1; j <= len(p); j++ {
		if p[j-1] != '*' {
			break
		}
		dp[0][j] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
				// dp[i][j] = dp[i][j] || dp[i][j-1] // 匹配空
				// dp[i][j] = dp[i][j] || dp[i-1][j] // 匹配任意不为空字符
			}
		}
	}
	return dp[len(s)][len(p)]
}

func isMatch(s string, p string) bool {
	// 这里是难点，字符串匹配都要加这个，相当于链表的虚拟结点
	s = " " + s
	p = " " + p
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				for k := 0; i-k >= 0; k++ {
					/*
						p[j] 为 '*'：可匹配任意长度的字符，可以匹配 0 个字符、匹配 1 个字符、匹配 2 个字符
						3.1. 当匹配为 0 个：f(i,j) = f(i, j - 1)
						3.2. 当匹配为 1 个：f(i,j) = f(i - 1, j - 1)
						3.3. 当匹配为 2 个：f(i,j) = f(i - 2, j - 1)
					*/
					// dp[i][j] = dp[i][j] || dp[i-k][j]
					dp[i][j] = dp[i-k][j-1] || dp[i][j]
					if dp[i][j] {
						break
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
