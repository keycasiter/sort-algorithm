package bubble

import (
	"fmt"
	"testing"
)

//递归
func TestBubbleSortRecurse(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}
	bubbleSortRecurse(nums, len(nums))
	//打印
	fmt.Println(nums)
}

func bubbleSortRecurse(nums []int, size int) {
	//不满足两个不需要比较
	if size < 2 {
		return
	}
	//这里for循环代表每一趟比较、交换的过程
	//每次都是从[0,n]的范围，只不过n会每次由外循环传递进来，并且每一趟最后一个会排序好，所以n的数量会递减
	//j代表每次遍历的下标，j<i-1 代表只遍历到数据规模的倒数第二个，因为要进行对比，这里目的是防止索引越界
	for j := 0; j < size-1; j++ {
		//nums[j]代表当前遍历到的元素，nums[j+1]代表下一个
		if nums[j+1] < nums[j] { //判断元素大小
			//交换过程
			tmp := nums[j+1]
			nums[j+1] = nums[j]
			nums[j] = tmp
		}
	}
	//每次数据规模减少1，因为每一次最靠后的都排序好了
	size--
	bubbleSortRecurse(nums, size)
}
