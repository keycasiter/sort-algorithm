package _select

import (
	"fmt"
	"testing"
)

// 递归 + 二元选择、剪枝优化
func TestSelectSortRecurseOptimize(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	selectSortRecurseOptimize(nums, 0)

	fmt.Println(nums)
}

func selectSortRecurseOptimize(nums []int, i int) {
	//i是数据规模开始的位置，i == len -1 代表当前只有一个元素，此时不需要排序了
	if i == len(nums)-1 {
		return
	}

	//每次遍历取最开头的元素作为游标对象进行对比，通过它来找到最大或者最小的那个元素
	min := nums[i]
	max := nums[i]
	//minIdx、maxIdx用来记录一次遍历后最大或最小元素的索引位置，为后面做交换使用的
	minIdx := i             //这里取本次数据范围的第一个为最小值，因为最小的会被放在第一个
	maxIdx := len(nums) - 1 //这里取最后一个为最大的对比，因为最大的会被放在最后一个
	//这里从i+1开始就可以，因为cmp就是i位置的元素，不用和自己对比了
	for j := i + 1; j < len(nums); j++ {
		if nums[j] < min { //对比元素大小, < 这里是找到最小的
			min = nums[j]
			minIdx = j
		}
		if nums[j] > max { //对比元素大小, > 这里是找到最大的
			max = nums[j]
			maxIdx = j
		}
	}

	//剪枝优化
	//最小位置的元素已经是最小，且最大位置的元素已经是最大，则说明剩余数据范围的元素都已经排序好了，可以停止遍历
	if minIdx == i && maxIdx == len(nums)-1 {
		return
	}

	//交换元素
	//此时的min、max是这轮遍历后最大或者最小的那个元素值
	//minIdx是这次遍历最小元素的索引，我妈要把它放到本次数据范围的最前面
	minTmp := nums[minIdx]
	nums[minIdx] = nums[i]
	nums[i] = minTmp
	//maxIdx是这次遍历最大元素的索引，我们要把它放到本次数据范围的最后面
	maxTmp := nums[maxIdx]
	nums[maxIdx] = nums[len(nums)-1]
	nums[len(nums)-1] = maxTmp
	//每一次遍历都找到最大或者最小的元素并将其放在了当前数据范围的最开头
	//所以每次i前面都是排序好的，下一次排序从i的下一个开始
	i++
	selectSortRecurseOptimize(nums, i)
}
