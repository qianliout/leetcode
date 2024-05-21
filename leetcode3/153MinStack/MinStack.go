package main

import (
	"fmt"
	"math"
)

func main() {
	ss := Constructor()
	ss.Push(-2)
	ss.Push(0)
	ss.Push(-3)
	fmt.Println(ss.GetMin())
	ss.Pop()
	fmt.Println(ss.Top())
	fmt.Println(ss.GetMin())
}

type MinStack struct {
	MinData int
	Data    []int
}

func Constructor() MinStack {
	return MinStack{
		MinData: math.MaxInt32,
		Data:    make([]int, 0),
	}

}

func (this *MinStack) Push(val int) {
	this.Data = append(this.Data, val)
	if this.MinData >= val {
		this.MinData = val
	}
}

func (this *MinStack) Pop() {
	if len(this.Data) == 0 {
		return
	}
	this.Data = this.Data[:len(this.Data)-1]
	this.MinData = math.MaxInt32
	for _, ch := range this.Data {
		if ch <= this.MinData {
			this.MinData = ch
		}
	}
}

func (this *MinStack) Top() int {
	pop := this.Data[len(this.Data)-1]
	return pop
}

func (this *MinStack) GetMin() int {
	return this.MinData
}
