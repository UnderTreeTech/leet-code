package array

// [1,2,3,4,5,6,7]
// 3
func Rotate(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}
	k = k % len(nums)

	size := len(nums)
	for i := 0; i < k; i++ {
		last := nums[size-1]
		for j := 0; j < size; j++ {
			tmp := nums[j]
			nums[j] = last
			last = tmp
		}
	}
}

func RotateReverse(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}
	k = k % len(nums)

	Reverse(nums)
	Reverse(nums[:k])
	Reverse(nums[k:])
}
