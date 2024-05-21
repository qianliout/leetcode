package main

import (
	"fmt"
	"math"
	"sort"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(findRadius([]int{1, 2, 3}, []int{2}))
}

// 第一步以每一个房屋为研究中心，然后遍历加热站（可以利用二分法节省效率）找到房屋的前一个加热站和后一个加热站。
// 接着通过比较房屋和前后加热站的距离比较房屋和谁更近。
//
// 对于所有房屋h(1),h(2),h(3)···h(n),每个房屋与最近的加热站的距离为dis(n),n为房间编号。即每个房屋距离加热站
// 的最短距离disList为：
// disList=[dis(1),dis(2),dis(3)···dis(n)]
// 因此：
// 当取disList中的最大值dis-max时。一定满足dis-max大于等于disList中的每一项。即dis-max一定能够辐射到每一个房屋。
// 为此dis-max即为问题的解。
// 二分查找
func findRadius2(houses []int, heaters []int) int {
	if len(heaters) == 0 || len(heaters) == 0 {
		return 0
	}
	sort.Ints(heaters)
	sort.Ints(houses)
	ans := 0
	n := len(heaters) - 1

	for i := 0; i < len(houses); i++ {
		left, right := 0, n
		for right-left > 1 {
			mid := left + (right-left)/2
			// 正好相等
			if houses[i] == heaters[mid] {
				left, right = mid, mid
				// 房子子加热站的右边
			} else if houses[i] > heaters[mid] {
				left = mid
			} else {
				right = mid
			}
		}
		mi := Min(AbsSub(houses[i], heaters[left]), AbsSub(houses[i], heaters[right]))
		if ans < mi {
			ans = mi
		}
	}
	return ans
}

func findRadius(houses []int, heaters []int) int {
	if len(heaters) == 0 || len(heaters) == 0 {
		return 0
	}
	sort.Ints(heaters)
	sort.Ints(houses)

	le, ri := 0, int(math.Pow10(9))

	for le < ri {
		mid := le + (ri-le)>>1
		if check(houses, heaters, mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

func check2(houses []int, heaters []int, dis int) bool {
	n, m := len(houses), len(heaters)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if houses[i] > heaters[j]+dis && (j+1 < m && houses[i] < heaters[j+1]-dis) {
				return false
			}
		}
	}
	return true
}

// 这个方法不好理解
func check(houses []int, heaters []int, x int) bool {
	n, m := len(houses), len(heaters)
	for i, j := 0, 0; i < n; i++ {
		for j < m && houses[i] > heaters[j]+x {
			j++
		}
		if j < m && heaters[j]-x <= houses[i] && houses[i] <= heaters[j]+x {
			continue
		}
		return false
	}
	return true
}
