package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home/home//home"))
}

func simplifyPath(path string) string {
	ss := strings.Split(path, "/")
	sta := make([]string, 0)
	for i := range ss {
		if ss[i] == "" || ss[i] == "." {
			continue
		}
		if ss[i] == ".." {
			if len(sta) > 0 {
				sta = sta[:len(sta)-1]
			}
			continue
		}
		sta = append(sta, ss[i])
	}

	p := strings.Join(sta, "/")
	return "/" + p
}
