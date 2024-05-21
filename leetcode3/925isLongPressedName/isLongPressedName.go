package main

import (
	"fmt"
)

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
	fmt.Println(isLongPressedName("alex", "aaleexa"))
	fmt.Println(isLongPressedName("vtkgn", "vttkgnn"))
}

func isLongPressedName(name string, typed string) bool {
	i := 0
	for j := 0; j < len(typed); j++ {
		if i < len(name) && name[i] == typed[j] {
			i++
		} else if i-1 >= 0 && i-1 < len(name) && typed[j] == name[i-1] {
			continue
		} else {
			return false
		}
	}
	return i > len(name)-1
}
