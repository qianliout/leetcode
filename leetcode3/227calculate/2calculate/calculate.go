package main

import (
	"fmt"
)

func main() {
	// fmt.Println(calculate("3+2*2"))
	// fmt.Println(calculate("3+2/2"))
	// fmt.Println(calculate(" 2/2 "))
	// fmt.Println(calculate(" 3/2 "))
	// fmt.Println(calculate(" (-5*3) "))
	// fmt.Println(calculate(" -5-3 "))
	fmt.Println(calculate("-5*-3"))
	// fmt.Println(calculate(" (1+(4+5)+(2)-3)+(6)+(8) "))
}

// -5*-3 这种会有 bug
// 表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
func calculate(s string) int {
	// s = strings.ReplaceAll(s, " ", "")
	ss := []byte(s)
	nums := make([]int, 0)
	nums = append(nums, 0)
	opsMap := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	ops := make([]byte, 0)
	// 为防止 () 内出现的首个字符为运算符，将所有的空格去掉，并将 (- 替换为 (0-，(+ 替换为 (0+（当然也可以不进行这样的预处理，将这个处理逻辑放到循环里去做）
	idx := 0
	for idx < len(ss) {
		ch := ss[idx]
		if ch == ' ' {
			idx++
			continue
		}
		if ch == '(' {
			ops = append(ops, ch)
			idx++
			continue
		}
		if ch == ')' {
			// 计算到最近一个左括号为止
			for len(ops) > 0 {
				if ops[len(ops)-1] != '(' {
					calc(&nums, &ops)
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
			idx++
			continue
		}

		if ch >= '0' && ch <= '9' {
			// 找到所有的数字
			j := idx
			nu := 0
			for j < len(ss) {
				if ss[j] > '9' || ss[j] < '0' {
					break
				}
				nu = nu*10 + int(ss[j]-'0')
				j++
			}
			nums = append(nums, nu)
			idx = j
			continue
		}

		// 走到这里就是加减乘除等符号了（不包括括号）
		if idx > 0 && !digit(ss[idx-1]) {
			// (-5*3), (+5+2) 这种，括号后面加上负号
			// 为防止 () 内出现的首个字符为运算符，将所有的空格去掉，并将 (- 替换为 (0-，(+ 替换为 (0+（当然也可以不进行这样的预处理，将这个处理逻辑放到循环里去做）
			nums = append(nums, 0)
		}

		// 有一个新操作要入栈时，先把栈内可以算的都算了
		// 只有满足「栈内运算符」比「当前运算符」优先级高/同等，才进行运算
		for len(ops) > 0 && ops[len(ops)-1] != '(' {
			prev := ops[len(ops)-1]
			if opsMap[prev] >= opsMap[ch] {
				calc(&nums, &ops)
			} else {
				break
			}
		}
		idx++
		ops = append(ops, ch)
	}

	for len(ops) > 0 {
		// fmt.Println("pre", nums, ops)
		calc(&nums, &ops)
	}
	// fmt.Println("end ", nums)
	return nums[len(nums)-1]
}

// 因为会改变大小，所以一定要传指针
// 如果不传指针则需要用返回值
func calc(nums *[]int, ops *[]byte) {
	fmt.Println("pre calc", *nums, *ops)
	if len(*nums) < 2 || len(*ops) == 0 {
		return
	}
	n := len(*nums)
	a, b := (*nums)[n-2], (*nums)[n-1]
	*nums = (*nums)[:n-2]
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]
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
	*nums = append(*nums, ans)
	fmt.Println("after calc", *nums, *ops)
}

func digit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
