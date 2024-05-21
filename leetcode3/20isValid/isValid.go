package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	ss := []byte(s)
	mp := map[byte]byte{'}': '{', ']': '[', ')': '('}

	for _, ch := range ss {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != mp[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
