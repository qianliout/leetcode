package main

import (
	"fmt"
)

func main() {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	fmt.Println(palindromePairs(words))
}

func palindromePairs(words []string) [][]int {
	ans := make([][]int, 0)
	mem := make(map[string]int)
	for i := range words {
		mem[words[i]] = i + 1 // 防止0值
	}
	for i := range words {
		word := []byte(words[i])
		for j := 0; j <= len(word); j++ {
			prefix := word[:j]
			postfix := word[j:]
			prefixR := reverse(prefix)
			// 当word的前缀的反转在字典中，且不是word自身，且word剩下部分是回文(空也是回文)
			// 则说明存在能与word组成回文的字符串
			if mem[prefixR] != 0 && mem[prefixR] != i+1 && palindrome(postfix) {
				ans = append(ans, []int{i, mem[prefixR] - 1})
			}
			// 当word的后缀的反转在字典中，且不是word自身，且word剩下部分是回文(空也是回文)
			// 则说明存在能与word组成回文的字符串

			// 为啥 j==0 要过滤了呢，因为会重复
			// 比如 abcd,当j=4时，prefix="abcd",postfix="",当j=0时，prefix="",postfix="abcd"
			// 这样就就会重复答案,所以前缀取了"",后缀判断时就不能取""了
			if j == 0 {
				continue
			}
			postfixR := reverse(postfix)
			if mem[postfixR] != 0 && mem[postfixR] != i+1 && palindrome(prefix) {
				ans = append(ans, []int{mem[postfixR] - 1, i})
			}
		}
	}
	return ans
}

func reverse(data []byte) string {
	ans := make([]byte, len(data))
	copy(ans, data)
	l, r := 0, len(data)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return string(ans)
}

func palindrome(data []byte) bool {
	l, r := 0, len(data)-1
	for l < r {
		if data[l] != data[r] {
			return false
		}
		l++
		r--
	}
	return true
}
