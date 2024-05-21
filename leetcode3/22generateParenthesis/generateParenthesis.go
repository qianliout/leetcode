package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis1(n int) []string {
	path := make([]byte, 0)
	ans := make([]string, 0)
	dfs(n, 0, 0, path, &ans)
	return ans
}

func dfs(n, left, right int, path []byte, ans *[]string) {
	if left == right && right == n {
		*ans = append(*ans, string(path))
		return
	}
	if left < n {
		path = append(path, '(')
		dfs(n, left+1, right, path, ans)
		path = path[:len(path)-1]
	}
	if right < left && right < n {
		path = append(path, ')')
		dfs(n, left, right+1, path, ans)
		path = path[:len(path)-1]
	}
}

func generateParenthesis2(n int) []string {
	if n <= 0 {
		return []string{""}
	}
	res := make([]string, 0)
	for i := 0; i < n; i++ {
		left := generateParenthesis(i)
		right := generateParenthesis(n - i - 1)
		for _, le := range left {
			for _, ri := range right {
				ans := "(" + le + ")" + ri
				res = append(res, ans)
			}
		}
	}
	return res
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	if n == 0 {
		ans = append(ans, "")
		return ans
	}

	for i := 0; i < n; i++ {
		left := generateParenthesis2(i)
		right := generateParenthesis2(n - i - 1)
		for _, le := range left {
			for _, ri := range right {
				// ans = append(ans, "("+le+")"+ri)
				ans = append(ans, "("+le+ri+")")
			}
		}
	}
	return ans
}
