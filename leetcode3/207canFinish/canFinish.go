package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
	fmt.Println(canFinish(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// in 入度，表示A 依赖的其他课程
	// out 出度，表示被 A 依赖的课程
	in := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		in[i] = make([]int, 0)
	}

	for i := range prerequisites {
		per := prerequisites[i]
		// 先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
		in[per[0]] = append(in[per[0]], per[1])
	}
	// 找出没有依赖项的，就可以先学习了
	for {
		del := make([]int, 0)
		for i, ch := range in {
			if len(ch) == 0 {
				del = append(del, i)
			}
		}
		if len(del) == 0 {
			break
		}
		// 可以先学习这些课程
		for i := range del {
			delete(in, del[i])
		}
		numCourses -= len(del)
		// 依赖这些课程的其他课程已经学习过了，所以依赖项目减少
		for _, study := range del {
			for i, depend := range in {
				in[i] = reduce(depend, study)
			}
		}
	}
	// 以下两种返回都是正确的
	// return len(in) == 0
	return numCourses == 0
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
