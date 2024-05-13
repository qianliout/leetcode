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

func dfs(nums []byte, start int, path string, cur int, pre int, target int, ans *[]string) {
	if start >= len(nums) {
		if cur == target {
			*ans = append(*ans, path)
		}
		return
	}
	for i := start; i < len(nums); i++ {
		// 不能有 02 这 样的数字
		if i > start && nums[start] == '0' {
			break
		}
		pa := string(nums[start : i+1])
		num, _ := strconv.Atoi(pa)

		if start == 0 {
			dfs(nums, i+1, pa, num, num, target, ans)
		} else {
			dfs(nums, i+1, path+"+"+pa, cur+num, num, target, ans)
			dfs(nums, i+1, path+"-"+pa, cur-num, -num, target, ans)
			// 这里的难点是记录 pre 的值
			// 例如 3+4*5*6,计算到5时，pre值是4，但是计算到6时，如果 pre值取5就会出错（结合,cur-pre+pre*num思考）
			// 所以 pre 值保存的是可以和*联合运算的值
			dfs(nums, i+1, path+"*"+pa, cur-pre+pre*num, num*pre, target, ans)
		}
	}
}
