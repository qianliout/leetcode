package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
}

func findMinHeightTrees(n int, edges [][]int) []int {
	inp := make(map[int][]int)
	for i := 0; i < n; i++ {
		inp[i] = make([]int, 0)
	}
	for i := range edges {
		fir, sec := edges[i][0], edges[i][1]
		inp[fir] = append(inp[fir], sec)
		inp[sec] = append(inp[sec], fir)
	}

	for len(inp) > 2 {
		lev := make([][2]int, 0)
		for i, ch := range inp {
			if len(ch) == 1 {
				lev = append(lev, [2]int{i, ch[0]})
			}
		}
		if len(lev) == 0 {
			break
		}

		// for _, ch := range lev {
		// 	delete(inp, ch[0])
		// }

		for _, ch := range lev {
			fir, sec := ch[0], ch[1]
			delete(inp, fir)

			inp[sec] = reduce(inp[sec], fir)
		}
	}

	ans := make([]int, 0)

	for i := range inp {
		ans = append(ans, i)
	}
	return ans
}

func reduce(aa []int, val int) []int {
	for i, ch := range aa {
		if ch == val {
			aa = append(aa[:i], aa[i+1:]...)
			break
		}
	}
	return aa
}
