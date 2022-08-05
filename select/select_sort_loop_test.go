package _select

import (
	"fmt"
	"testing"
)

// 遍历
func TestSelectSortLoop(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	//这里的for循环是进行n次遍历，每一次的目的是把当前数据规模中最大或者最小的元素找到
	//然后将找到（选择）的这个元素放到当前数据范围的最前面，这样随着遍历的推移，每一次都会找到当前数据规模最大或者最小的元素
	//靠前面的数据已经排好序，靠后面的就是还没被排序的部分
	for i := 0; i < len(nums); i++ {
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
	}
	fmt.Println(nums)
}
