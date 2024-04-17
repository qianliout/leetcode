package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("((((((((((((((((((((aaaaa"))
}

func removeInvalidParentheses(s string) []string {
	ss := []byte(s)
	ans := make([]string, 0)
	exit := make(map[string]bool)
	dfs(ss, 0, &ans)
	// 查最大的
	res, length := make([]string, 0), 0
	for i := range ans {
		if len(ans[i]) > length {
			length = len(ans[i])
		}
	}

	for i := range ans {
		if len(ans[i]) == length && !exit[ans[i]] {
			res = append(res, ans[i])
			exit[ans[i]] = true
		}
	}

	return res
}

func dfs(ss []byte, start int, ans *[]string) {
	// 这里的判断是一个出错点，这里不能放到 start>=len(ss)里面去判断
	if check(ss) {
		*ans = append(*ans, String(ss))
	}
	if start >= len(ss) {
		return
	}
	for i := start; i < len(ss); i++ {
		// 因为有其他字符串
		if ss[i] != '(' && ss[i] != ')' && ss[i] != '@' {
			continue
		}
		pre := ss[i]
		ss[i] = '@'
		dfs(ss, i+1, ans)
		ss[i] = pre
	}
}

func check(data []byte) bool {
	stark := make([]byte, 0)
	for _, ch := range data {
		if ch != '(' && ch != ')' {
			continue
		}
		if ch == '(' {
			stark = append(stark, '(')
			continue
		}
		if len(stark) == 0 {
			return false
		}
		stark = stark[:len(stark)-1]
	}

	return len(stark) == 0
}

func String(data []byte) string {
	res := make([]byte, 0)
	for _, ch := range data {
		if ch != '@' {
			res = append(res, ch)
		}
	}
	return string(res)
}
