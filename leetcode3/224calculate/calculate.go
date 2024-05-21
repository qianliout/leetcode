package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate("1+1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("2-(1 + 2)"))
}

func calculate(s string) int {
	ans, num, op := 0, 0, 1
	nstak := make([]int, 0)
	sstak := make([]int, 0)
	for _, ch := range s {
		switch ch {
		case ' ':
			continue
		case '+':
			ans = ans + num*op
			num, op = 0, 1
		case '-':
			ans = ans + num*op
			num, op = 0, -1
		case '(':
			ans = ans + num*op
			nstak = append(nstak, ans)
			sstak = append(sstak, op)
			ans, num, op = 0, 0, 1
		case ')':
			ans = ans + num*op
			ans = ans * sstak[len(sstak)-1]
			ans = ans + nstak[len(nstak)-1]
			sstak = sstak[:len(sstak)-1]
			nstak = nstak[:len(nstak)-1]
			num, op = 0, 1
		default:
			num = num*10 + int(ch) - 48
		}
	}
	return ans + op*num
}
