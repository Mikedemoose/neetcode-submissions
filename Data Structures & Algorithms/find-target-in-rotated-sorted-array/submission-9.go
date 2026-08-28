func search(nums []int, target int) int {
	return rotatedBinarySearch(nums, 0, len(nums)-1, target)
}

func rotatedBinarySearch(nums []int, left, right int, target int) int {
	if len(nums) == 0 || left > right {
		return -1
	}
	if left == right && target != nums[left] {
		return -1
	}
	// case 1: full array is sorted
	if nums[left] < nums[right] {
		return standardBinarySearch(nums, left, right, target)
	}

	mid := (left+right)/2
	if nums[mid] == target {
		return mid
	}
	// case 2: only first half is sorted
	if nums[left] < nums[mid] {
		if target >= nums[left] && target < nums[mid] {
			return standardBinarySearch(nums, left, mid-1, target)
		}
		return rotatedBinarySearch(nums, mid+1, right, target)
	}
	// case 3: only second half is sorted
	// will be true by this stage
	if target >= nums[mid+1] && target <= nums[right] {
		return standardBinarySearch(nums, mid+1, right, target)
	}
	return rotatedBinarySearch(nums, left, mid-1, target)
}





func standardBinarySearch(nums []int, left, right int, target int) int {
	if len(nums) == 0 || left > right {
		return -1
	}

	for left <= right {
		mid := (left+right)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return -1
}
