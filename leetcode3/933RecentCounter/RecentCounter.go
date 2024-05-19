package main

func main() {

}

type RecentCounter struct {
	Data []int
}

func Constructor() RecentCounter {
	return RecentCounter{Data: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.Data) > 0 && this.Data[0] < t-3000 {
		this.Data = this.Data[1:]
	}
	this.Data = append(this.Data, t)
	return len(this.Data)
}
