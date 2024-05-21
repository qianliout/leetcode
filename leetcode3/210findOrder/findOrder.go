package main

import (
	"fmt"
)

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	// in 入度，表示A 依赖的其他课程
	in := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		in[i] = make([]int, 0)
	}
	// 课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
	for _, ch := range prerequisites {
		in[ch[0]] = append(in[ch[0]], ch[1])
	}
	ans := make([]int, 0)
	for {
		cou := make([]int, 0)
		for i, ch := range in {
			if len(ch) == 0 {
				cou = append(cou, i)
			}
		}
		if len(cou) == 0 {
			break
		}
		ans = append(ans, cou...)
		for _, ch := range cou {
			delete(in, ch)
		}
		// 依赖这些课程的其他课程已经学习过了，所以依赖项目减少
		for _, study := range cou {
			for i, depend := range in {
				in[i] = reduce(depend, study)
			}
		}
	}
	if len(in) > 0 {
		return []int{}
	}
	return ans
}

func reduce(aa []int, val int) []int {
	res := make([]int, 0)
	for _, ch := range aa {
		if ch != val {
			res = append(res, ch)
		}
	}
	return res
}
