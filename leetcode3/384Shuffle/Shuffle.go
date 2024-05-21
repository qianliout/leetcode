package main

import (
	"math/rand"
	"time"
)

func main() {

}

type Solution struct {
	Data []int
	Rand *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		Data: nums,
		Rand: rand.New(rand.NewSource(time.Now().UnixMicro())),
	}
}

func (this *Solution) Reset() []int {
	this.Rand = rand.New(rand.NewSource(time.Now().UnixMicro()))
	return this.Data
}

func (this *Solution) Shuffle2() []int {
	nums := append([]int{}, this.Data...)
	for i := 0; i < len(nums); i++ {
		ra := this.Rand.Intn(len(this.Data) - i)
		nums[i], nums[i+ra] = nums[i+ra], nums[i]
	}
	return nums
}

// 推荐这种写法
func (this *Solution) Shuffle() []int {
	nums := append([]int{}, this.Data...)
	for i := len(nums) - 1; i >= 0; i-- {
		ra := this.Rand.Intn(i + 1)
		nums[i], nums[ra] = nums[ra], nums[i]
	}
	return nums
}
