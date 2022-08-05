### 核心点
**找到基准值的位置**
### 步骤
- 第一步，选择一个值作为基准值；
- 第二步，找到基准值的位置，并将小于基准值的元素放在基准值的前面，大于基准值的元素放在基准值的后面；
- 第三步，对基准值的左右两侧递归地进行这个过程。
> 挖坑填数 + 分治

### 实现方式

```go

func TestQuickSort1(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

// nums数据元素
// left左下标
// right右下标
func quickSort(nums []int, left, right int) {
	// 左下标大于等于右下标退出，此时数据规模不再需要排序
	if left >= right {
		return
	}
	//基准值，这里设置为取每次数据规模的左下标起始位置的元素作为中轴线
	pivot := nums[left]
	//左下标
	leftIdx := left
	//右下标
	rightIdx := right

	//只要左下标小于右下标就运行，等于时说明左右下标重合，此时不需要执行
	for leftIdx < rightIdx {
		//移动左
		for nums[leftIdx] < pivot {
			leftIdx++
		}
		for nums[rightIdx] > pivot {
			rightIdx--
		}
		nums[leftIdx], nums[rightIdx] = nums[rightIdx], nums[leftIdx]
	}
	//此时循环结束，左右下标重合，将基准值放入
	nums[leftIdx] = pivot
	quickSort(nums, left, leftIdx-1)
	quickSort(nums, rightIdx+1, right)
}
```

### 优化
- 采用更合理的**基准数(中心轴)**，减少递归的深度，从数列中选取多个数，取中间数
- 结合插入排序，区间在10个元素之内采用插入排序，效率更高