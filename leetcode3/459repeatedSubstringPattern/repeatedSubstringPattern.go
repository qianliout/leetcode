package main

import (
	"strings"
)

func main() {

}

func repeatedSubstringPattern(s string) bool {
	ss := []byte(s)
	for i := 1; i < len(ss); i++ {
		if strings.Repeat(string(ss[:i]), len(s)/i) == s {
			return true
		}
	}
	return false
}
