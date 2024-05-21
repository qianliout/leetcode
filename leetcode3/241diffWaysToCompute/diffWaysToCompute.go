package main

import (
	"strconv"
)

func main() {

}

func diffWaysToCompute(expression string) []int {
	return dfs([]byte(expression))
}

func dfs(ex []byte) []int {
	if n, err := strconv.Atoi(string(ex)); err == nil {
		return []int{n}
	}
	// 其实可以不用写这一步，因为上面的转化会包含这一步
	if len(ex) == 0 {
		return []int{0}
	}
	ans := make([]int, 0)
	for i := 0; i < len(ex); i++ {
		if ex[i] == '+' || ex[i] == '-' || ex[i] == '*' {
			left := dfs(ex[:i])
			right := dfs(ex[i+1:]) // 因为是合法的表达式，所以不会越界
			for _, l := range left {
				for _, r := range right {
					switch ex[i] {
					case '+':
						ans = append(ans, l+r)
					case '-':
						ans = append(ans, l-r)
					case '*':
						ans = append(ans, l*r)
					}
				}
			}
		}
	}
	return ans
}
