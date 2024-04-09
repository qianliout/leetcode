package main

import (
	"strings"
)

func main() {

}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	ss1 := make([]string, 0)

	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == "" {
			continue
		}
		ss1 = append(ss1, ss[i])

	}
	ans := strings.Join(ss1, " ")
	return strings.TrimSpace(ans)
}
