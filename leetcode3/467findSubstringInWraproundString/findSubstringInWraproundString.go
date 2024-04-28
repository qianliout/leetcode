package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	// fmt.Println(findSubstringInWraproundString("zab"))
	// fmt.Println(findSubstringInWraproundString("cac"))
	// fmt.Println(findSubstringInWraproundString("cdefghefghijklmnopqrstuvwxmnijklmnopqrstuvbcdefghijklmnopqrstuvwabcddefghijklfghijklmabcdefghijklmnopqrstuvwxymnopqrstuvwxyz"))
	// fmt.Println(findSubstringInWraproundString2("cdefghefghijklmnopqrstuvwxmnijklmnopqrstuvbcdefghijklmnopqrstuvwabcddefghijklfghijklmabcdefghijklmnopqrstuvwxymnopqrstuvwxyz"))
	fmt.Println(findSubstringInWraproundString("zaba"))
}

func findSubstringInWraproundString2(s string) int {
	dp := make([]int, 26)
	/*
		因为s是一个固定的字符串，我们维护以每个字符结尾的最长子串长度，就可以直接累加所有以该字符结尾的子字符串 (每个字符统计最长长度即可。因为更小的长度都是更长的长度的子串)。
		在遍历的时候我们维护一个当前的最长长度。
		如果当前字符满足与上一个字符在s中连续（也就是后面的比前面的ASCII的差模26大1），那么以当前字符结尾的最长长度就从前面累加，
		否则就是新的开始。
	*/

	cnt := 1
	for i := range s {
		idx := s[i] - 'a'
		if i > 0 {
			if check(s[i-1], s[i]) {
				cnt++
			} else {
				cnt = 1
			}
		}

		if dp[idx] < cnt {
			dp[idx] = cnt
		}
	}

	ans := 0
	for _, ch := range dp {
		ans += ch
	}
	return ans
}

func findSubstringInWraproundString(s string) int {
	// dp[i]是指s到i这个位置（包括i），存在的最长的子串位数
	// 情况一：在base 中，s[i]是a[i-1]的下一个字符，说明接的上，dp[i] = dp[i-1]+1
	// 情况二：在base 中，s[i]不是s[i-1]的下一个字符，说明接不上，dp[i] = +1

	// 要去重
	// 初值 dp[i]=1
	dp := make(map[int]int)
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		dp[i] = Max(1, dp[idx])
		if i > 0 {
			if check(s[i-1], s[i]) {
				dp[i] = Max(dp[i], dp[i-1]+1)
			} else {
				dp[i] = 1
			}
		}
	}
	ans := 0
	for _, ch := range dp {
		ans += ch
	}
	return ans
}

func check(pre, after byte) bool {
	if after-pre == 1 {
		return true
	}
	if pre == 'z' && after == 'a' {
		return true
	}
	return false
}
