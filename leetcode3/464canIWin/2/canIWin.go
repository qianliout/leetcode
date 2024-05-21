package main

import (
	"fmt"
)

func main() {
	// fmt.Println(canIWin(10, 11)) // false
	// fmt.Println(canIWin(10, 0))  // true
	// fmt.Println(canIWin(10, 1))  // true
	// fmt.Println(canIWin(18, 79)) // true
	// fmt.Println(canIWin(10, 40)) // false
	fmt.Println(canIWin(5, 50)) // false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	/*
		当把 1 到 maxChoosableInteger所有的数字全都累加，依然无法到达 desiredTotal， 先手一定输了。
		为啥？因为我们问的是先手能不能赢，假如总和都达不到 desiredTotal，先手怎么赢？
		另外一种情况是 maxChoosableInteger>=desiredTotal 那么先手一定赢。
	*/
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	mem := make(map[int]bool)
	return dfs2(0, 0, desiredTotal, maxChoosableInteger, mem)

}

// timeout
// dfs的结果和两个值有关，sum和choose，又因为 choose 是一个 map,不好缓存
func dfs(choose map[int]bool, sum int, desired int) bool {
	for k, v := range choose {
		if v {
			continue
		}
		if sum+k >= desired {
			return true
		}
		choose[k] = true
		other := dfs(choose, sum+k, desired)
		/*
			// 这种方法不可以，因为不管结果如何，都应该做回退
			if !dfs(choose,sum+k,desired){
				return true
			}
		*/

		choose[k] = false
		if !other {
			return true
		}
	}
	return false
}

// 候选集最大二，所以可以用二进制优化
// 然后进一步可以加缓存
// 为啥 mem中只缓存了 stat
// 同时，只要知道 stat，就能知道「已选择的数字」有哪些，因此也就确定了当前的 sum。
// 所以，只需要知道 stat，就能确定 dfs()的输出。
func dfs2(stat int, sum int, desired int, maxC int, mem map[int]bool) bool {
	if v, ok := mem[stat]; ok {
		return v
	}

	for k := 1; k <= maxC; k++ {
		// 这一位是1，说明前面已经选择了
		if stat&(1<<k) > 0 {
			continue
		}
		if sum+k >= desired {
			mem[stat] = true
			return true
		}
		// 把第i位设置成1
		// stat = stat | (1 << k)
		other := dfs2(stat|(1<<k), sum+k, desired, maxC, mem)
		// 这里有个技巧,把第 i 位设置成0，是不好做的，那么就把设置值的这个动作放到下一层
		// 把 第i位设置成0
		// stat = stat & (^(1 << k))
		if !other {
			mem[stat] = true
			return true
		}
	}

	mem[stat] = false
	return false
}
