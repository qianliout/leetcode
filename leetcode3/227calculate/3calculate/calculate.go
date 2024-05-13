package main

import (
	"fmt"
)

func main() {
	// fmt.Println(calculate("1+1"))
	fmt.Println(calculate("3/2"))
	// fmt.Println(calculate("0-2147483647"))
}

// -5*-3 这种会有 bug
// 表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
func calculate(s string) int {
	ss := []byte(s)

	nums := make([]int, 0)
	var preOP byte = '+'

	ans := 0
	num := 0
	// 这样写 每次的 i 都是新的变量，所以下面对 i 的操作是不会生效的
	// for i := range ss {
	for i := 0; i < len(ss); i++ {
		if ss[i] == ' ' {
			continue
		}
		for i < len(s) {
			ch := s[i]
			if ch >= '0' && ch <= '9' {
				num = num*10 + int(ch) - 48
				i++
			} else {
				break
			}
		}
		switch preOP {
		case '+':
			nums = append(nums, num)
		case '-':
			nums = append(nums, -num)
		case '*':
			if len(nums) == 0 {
				return 0
			}
			last := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums = append(nums, last*num)
		case '/':
			if len(nums) == 0 || num == 0 {
				return 0
			}
			last := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			nums = append(nums, last/num)
		}
		// 这里容易出错
		if i < len(ss) {
			num, preOP = 0, byte(ss[i])
		}
	}

	for _, nu := range nums {
		ans += nu
	}
	return ans
}

// 因为会改变大小，所以一定要传指针
// 如果不传指针则需要用返回值
func calc(nums []int, ops []byte) ([]int, []byte) {
	if len(nums) < 2 || len(ops) == 0 {
		return nums, ops
	}
	n := len(nums)
	a, b := nums[n-2], nums[n-1]
	nums = nums[:n-2]
	op := ops[len(ops)-1]
	ops = ops[:len(ops)-1]
	ans := 0
	switch op {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	case '/':
		// 这里要判断
		ans = a / b
	case '%':
		ans = a % b
		//  还可以加其他操作
	}
	nums = append(nums, ans)
	return nums, ops
}

func digit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
