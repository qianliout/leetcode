package main

import (
	"sort"
	"strings"
)

func main() {

}

func groupAnagrams(strs []string) [][]string {
	exit := make(map[string][]string)
	for i := range strs {
		st := strings.Split(strs[i], "")
		sort.Strings(st)
		st1 := strings.Join(st, "")
		if exit[st1] == nil {
			exit[st1] = make([]string, 0)
		}
		exit[st1] = append(exit[st1], strs[i])
	}
	ans := make([][]string, 0)
	for i := range exit {
		ans = append(ans, exit[i])
	}
	return ans
}
