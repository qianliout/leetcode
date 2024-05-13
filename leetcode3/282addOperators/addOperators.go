package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addOperators("123", 6))
	fmt.Println(addOperators("1203", 6))
	fmt.Println(addOperators("3456237490", 9191))
}

func addOperators(num string, target int) []string {
	ans := make([]string, 0)

	dfs([]byte(num), 0, "", 0, 0, target, &ans)
	return ans
}

func dfs(nums []byte, start int, path string, pre int, cur, target int, res *[]string) {
	if start >= len(nums) {
		if cur == target {
			*res = append(*res, path)
		}
		return
	}

	for i := start; i < len(nums); i++ {
		// 中间的字符串 02，计算结是2,这种情况是不准的
		if nums[start] == '0' && i > start {
			break
		}
		nu, _ := strconv.Atoi(string(nums[start : i+1]))
		pa := string(nums[start : i+1])

		if start == 0 {
			dfs(nums, i+1, pa, nu, nu, target, res)
		} else {
			dfs(nums, i+1, path+"+"+pa, nu, cur+nu, target, res)
			dfs(nums, i+1, path+"-"+pa, -nu, cur-nu, target, res)
			dfs(nums, i+1, path+"*"+pa, pre*nu, cur-pre+pre*nu, target, res)
		}
	}
}
