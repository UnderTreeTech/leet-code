package array

func Reverse(nums []int) {
	if len(nums) < 2 {
		return
	}

	start := 0
	end := len(nums) - 1
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}
