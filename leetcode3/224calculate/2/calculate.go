package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate("1+1"))
}

func calculate(s string) int {
	ans, op, num := 0, 1, 0
	nums := make([]int, 0)
	ops := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}

		switch s[i] {
		case '+':
			ans = ans + num*op
			op, num = 1, 0
		case '-':
			ans = ans + num*op
			op, num = -1, 0
		case '(':
			ans = ans + num*op
			nums = append(nums, ans)
			ops = append(ops, op)
			ans, op, num = 0, 1, 0
		case ')':
			ans = ans + num*op
			ans = ans * ops[len(ops)-1]
			ans = ans + nums[len(nums)-1]
			ops = ops[:len(ops)-1]
			nums = nums[:len(nums)-1]
			num, op = 0, 1
		default:
			num = num*10 + int(s[i]-'0')
		}
	}
	return ans + op*num
}
