package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(findKthNumber(13, 2))
}

// 给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。
func findKthNumber(n int, k int) int {
	pre := 1
	// 为啥k>1呢，因为 pre=1已经占了一个数了
	for k > 1 {
		cnt := getCnt(pre, n)
		if cnt < k {
			pre++
			k = k - cnt
		} else {
			pre = pre * 10
			k--
		}
	}

	return pre
}

// 该函数实现了统计范围 [1,limit] 内以 prefix 为前缀的数的个数。
func getCnt(prefix, limit int) int {
	nexPrefix := prefix + 1
	curPrefix := prefix
	cnt := 0
	for curPrefix <= limit {
		// 这一步是难点
		cnt += Min(nexPrefix-curPrefix, limit-curPrefix+1)
		curPrefix = curPrefix * 10
		nexPrefix = nexPrefix * 10
	}
	return cnt
}
