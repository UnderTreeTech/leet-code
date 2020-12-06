package array

func RemoveSortedDupElements(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	var dup,gap int
	pivot := nums[0]
	for i:=1; i<size;i++ {
		if pivot == nums[i] {
			dup++
			gap++
			continue
		}

		pivot = nums[i]
		if gap >0 {
			nums[i-gap] = nums[i]
		}
	}

	return size - dup
}
