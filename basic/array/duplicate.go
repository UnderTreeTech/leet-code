package array

// 删除有序数组中重复元素
// 非有序数组可以先排序，排序算法
func RemoveSortedDupElements(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	var dup, gap int
	pivot := nums[0]
	for i := 1; i < size; i++ {
		if pivot == nums[i] {
			dup++
			gap++
			continue
		}

		pivot = nums[i]
		if gap > 0 {
			nums[i-gap] = nums[i]
		}
	}

	return size - dup
}
