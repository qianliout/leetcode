package main

func main() {
}

// type Entry struct {
// 	keys map[string]bool
// 	freq int
// }
//
// type AllOne struct {
// 	Data map[string]*list.Element
// 	Head *list.List
// }
//
// func Constructor() AllOne {
// 	return AllOne{
// 		Head: list.New(),
// 	}
// }
//
// func (vi *AllOne) Inc(key string) {
// 	el := vi.Data[key]
// 	if el == nil {
//
// 		return
// 	}
//
// }
//
// func (vi *AllOne) Dec(key string) {
// 	if _, ok := vi.Data[key]; !ok {
// 		return
// 	}
// 	vi.Data[key]--
// 	if vi.Data[key] < vi.MinCnt {
// 		vi.Min = key
// 		vi.MinCnt = vi.Data[key]
// 	}
// }
//
// func (vi *AllOne) GetMaxKey() string {
// 	return vi.Max
//
// }
//
// func (vi *AllOne) GetMinKey() string {
// 	return vi.Min
// }
