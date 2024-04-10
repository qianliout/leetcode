package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate("3+2*2"))
}

// 只有加减乘除
func calculate(s string) int {
	ans, num := 0, 0
	nstak := make([]int, 0)
	var op byte = '+'

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
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

		switch op {
		case '+':
			nstak = append(nstak, num)
		case '-':
			nstak = append(nstak, -num)
		case '*':
			last := nstak[len(nstak)-1]
			nstak = nstak[:len(nstak)-1]
			nstak = append(nstak, last*num)
		case '/':
			last := nstak[len(nstak)-1]
			nstak = nstak[:len(nstak)-1]
			nstak = append(nstak, last/num)
		}
		// 这里为啥会有这个判断呢，因为当最后一个数据里，i=len(s)
		if i < len(s) {
			num, op = 0, byte(s[i])
		}
	}
	for _, nu := range nstak {
		ans += nu
	}
	return ans
}
