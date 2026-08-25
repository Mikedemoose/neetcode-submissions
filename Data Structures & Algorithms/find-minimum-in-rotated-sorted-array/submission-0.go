func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	mid := (left + right + 1) / 2

	for left < right{
		mid = (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[(left+right)/2]
}