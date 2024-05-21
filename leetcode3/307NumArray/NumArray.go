package main

import (
	"fmt"
)

func main() {
	na := Constructor([]int{1, 3, 5})
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))

}

type NumArray struct {
	Data   []int
	Prefix []int
}

func Constructor(nums []int) NumArray {

	na := NumArray{
		Data:   nums,
		Prefix: make([]int, len(nums)+1),
	}
	if len(nums) == 0 {
		return na
	}

	for i := 1; i < len(na.Prefix); i++ {
		na.Prefix[i] = na.Prefix[i-1] + nums[i-1]
	}
	return na
}

func (this *NumArray) Update(index int, val int) {
	if index >= len(this.Data) || index < 0 {
		return
	}

	pre := this.Data[index]
	sub := val - pre
	this.Data[index] = val
	for i := index + 1; i < len(this.Prefix); i++ {
		this.Prefix[i] = this.Prefix[i] + sub
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Prefix[right+1] - this.Prefix[left]
}
