package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(0))
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(0))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.GetRandom())
}

type RandomizedSet struct {
	Data []int
	Exit map[int]int // val->idx
	Rand *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Data: make([]int, 0),
		Exit: make(map[int]int),
		Rand: rand.New(rand.NewSource(time.Now().UnixMilli())),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Exit[val]; ok {
		return false
	}
	this.Data = append(this.Data, val)
	this.Exit[val] = len(this.Data) - 1
	return true
}

// 可以实现，但是不是 O(1)
func (this *RandomizedSet) Remove1(val int) bool {
	idx, ok := this.Exit[val]
	if !ok {
		return false
	}

	delete(this.Exit, val)

	this.Data = append(this.Data[:idx], this.Data[idx+1:]...)

	for i, ch := range this.Data {
		this.Exit[ch] = i
	}

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.Exit[val]
	if !ok {
		return false
	}
	// 这里有个技巧，就是把要删除的元素和最后一个元素互换
	l := len(this.Data)

	this.Exit[this.Data[l-1]] = index
	this.Data[index], this.Data[l-1] = this.Data[l-1], this.Data[index]

	// this.Data[index] = this.Data[l-1]
	// this.Data[this.Data[index]] = index
	this.Data = this.Data[:l-1]
	delete(this.Exit, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.Data) == 0 {
		return 0
	}
	idx := this.Rand.Intn(len(this.Data))
	return this.Data[idx]
}
