package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 3, 3, 3, 5, 5, 5, 5}
	n := removeDuplicates(nums)
	fmt.Println(nums[:n])

}

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
说明：
为什么返回数值是整数，但输出的答案是数组呢？
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	left, i := 2, 2
	for i < len(nums) {
		if nums[i] == nums[left-1] && nums[i] == nums[left-2] {
			i++
		} else {
			nums[left] = nums[i]
			left++
			i++
		}
	}
	return left
}
