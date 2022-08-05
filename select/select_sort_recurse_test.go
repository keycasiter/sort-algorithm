package _select

import (
	"fmt"
	"testing"
)

// 递归
func TestSelectSortRecurse(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	selectSortRecurse(nums, 0)

	fmt.Println(nums)
}

func selectSortRecurse(nums []int, i int) {
	//i是数据规模开始的位置，i == len -1 代表当前只有一个元素，此时不需要排序了
	if i == len(nums)-1 {
		return
	}

	//每次遍历取最开头的元素作为游标对象进行对比，通过它来找到最大或者最小的那个元素
	cmp := nums[i]
	//idx用来记录一次遍历后最大或最小元素的索引位置，为后面做交换使用的
	idx := i
	//这里从i+1开始就可以，因为cmp就是i位置的元素，不用和自己对比了
	for j := i + 1; j < len(nums); j++ {
		if nums[j] < cmp { //对比元素大小, < 这里是找到最小的， > 的话是找到最大的
			cmp = nums[j]
			idx = j
		}
	}
	//交换元素
	//此时的cmp是这轮遍历后最大或者最小的那个元素值，idx是这个元素的索引，我们要把它放到本次数据范围的最前面
	tmp := nums[idx]
	nums[idx] = nums[i]
	nums[i] = tmp

	//每一次遍历都找到最大或者最小的元素并将其放在了当前数据范围的最开头
	//所以每次i前面都是排序好的，下一次排序从i的下一个开始
	i++
	selectSortRecurse(nums, i)
}
