package insert

import (
	"fmt"
	"testing"
)

//递归
func TestInsertSortRecurse(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}
	insertSortRecurse(nums, 1)
	fmt.Println(nums)
}

func insertSortRecurse(nums []int, i int) {
	//当i超过nums长度就停止
	if i <= 0 || i >= len(nums) {
		return
	}

	//j代表每一次遍历要对比的元素
	//由于上面提到 i 前面都是排序完成的，i后面都是未排序的
	//因此 j = i 是选取未排序的第一个元素来进行和i 前面未排序的元素依次对比完成本次排序
	for j := i; j > 0; j-- {
		if nums[j] < nums[j-1] { // > 把较小的往前移; < 把较大的往前移
			//交换元素
			tmp := nums[j]
			nums[j] = nums[j-1]
			nums[j-1] = tmp
		}
	}
	//每一轮排序完，i递增，从下一个继续开始
	i++
	insertSortRecurse(nums, i)
}
