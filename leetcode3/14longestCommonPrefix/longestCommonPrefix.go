package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"abcvd", "ab", "abc"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	maxS := ""
	for i := range strs {
		if len(strs[i]) >= len(maxS) {
			maxS = strs[i]		}
 :NERDTreeToggle
 }
	ss := []byte(maxS)
	start := 0
	finished := false
	for range maxS {
		if finished {
			break
		}
		for j := range strs {
			if start >= len(strs[j]) {
				finished = true
				break
			}
			if strs[j][start] != maxS[start] {
				finished = true
				break
			}
		}
		if !finished {
			start++
		}
	}

	return string(ss[:start])
}
