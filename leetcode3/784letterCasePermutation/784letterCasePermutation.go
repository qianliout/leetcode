package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}

func letterCasePermutation(s string) []string {
	queue := make([][]byte, 0)
	queue = append(queue, []byte(s))
	res := make([]string, 0)
	used := make(map[string]bool)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		ans := find(first)
		for i := range ans {
			if !used[ans[i]] {
				used[ans[i]] = true
				res = append(res, ans[i])
				queue = append(queue, []byte(ans[i]))
			}
		}
	}
	return res
}

func find(ss []byte) []string {
	res := make([]string, 0)
	res = append(res, string(ss))

	for i, ch := range ss {
		if ch >= 'A' && ch <= 'Z' {
			ss[i] = ss[i] + 32
			res = append(res, string(ss))
			ss[i] = ss[i] - 32
		}
		if ch >= 'a' && ch <= 'z' {
			ss[i] = ss[i] - 32
			res = append(res, string(ss))
			ss[i] = ss[i] + 32
		}
	}
	return res
}
