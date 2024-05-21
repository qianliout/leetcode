package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rc := Constructor()
	rc.Insert(0)
	rc.Insert(1)
	rc.Remove(0)
	rc.Insert(2)
	rc.Remove(1)
	fmt.Println(rc.GetRandom())
}

type IndexMap map[int]bool

func (vi IndexMap) GetRandomIndex() (int, bool) {
	for idx, k := range vi {
		if k {
			return idx, true
		}
	}

	return 0, false
}

type RandomizedCollection struct {
	Data []int
	Exit map[int]IndexMap // key->map[idx]=true
	Rand *rand.Rand
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		Data: make([]int, 0),
		Exit: make(map[int]IndexMap),
		Rand: rand.New(rand.NewSource(time.Now().UnixMilli())),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	_, exit := this.Exit[val].GetRandomIndex()

	this.Data = append(this.Data, val)

	if this.Exit[val] == nil || !exit {
		this.Exit[val] = make(IndexMap)
	}

	la := len(this.Data) - 1

	this.Exit[val][la] = true

	return !exit
}

func (this *RandomizedCollection) Remove(val int) bool {
	idx, ok := this.Exit[val].GetRandomIndex()
	if !ok {
		delete(this.Exit, val)
		return false
	}

	this.Exit[val][idx] = false

	lastDataIndex := len(this.Data) - 1

	lastData := this.Data[lastDataIndex]

	if idx == lastDataIndex {
		this.Data = this.Data[:lastDataIndex]
	} else {
		// 把最后一个换到这个位置
		this.Data[idx], this.Data[lastDataIndex] = this.Data[lastDataIndex], this.Data[idx]
		this.Data = this.Data[:lastDataIndex]
		this.Exit[lastData][idx] = true
		this.Exit[lastData][lastDataIndex] = false
	}
	return true
}

func (this *RandomizedCollection) GetRandom() int {
	if len(this.Data) == 0 {
		return 0
	}
	idx := this.Rand.Intn(len(this.Data))
	return this.Data[idx]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
