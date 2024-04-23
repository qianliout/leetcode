package main

func main() {
	all := Constructor()
	all.Inc("hello")
	all.Inc("hello")
}

type AllOne struct {
	Data   map[string]int
	Max    string
	Min    string
	MaxCnt int
	MinCnt int
}

func Constructor() AllOne {
	return AllOne{
		Data:   make(map[string]int),
		Max:    "",
		Min:    "",
		MaxCnt: 0,
		MinCnt: 0,
	}
}

func (this *AllOne) Inc(key string) {
	this.Data[key]++
	if this.Data[key] > this.MinCnt {
		this.Max = key
		this.MaxCnt = this.Data[key]
	}
	if this.Min == "" {
		this.Max = key
		this.MaxCnt = this.Data[key]
	}
}

func (this *AllOne) Dec(key string) {
	if _, ok := this.Data[key]; !ok {
		return
	}
	this.Data[key]--
	if this.Data[key] < this.MinCnt {
		this.Min = key
		this.MinCnt = this.Data[key]
	}
}

func (this *AllOne) GetMaxKey() string {
	return this.Max

}

func (this *AllOne) GetMinKey() string {
	return this.Min
}
