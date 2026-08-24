func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxProduct := nums[0]
 	
	// go the first non-zero element
	i := 0
	for i<len(nums)-1 {
		if nums[i] > maxProduct {
			maxProduct = nums[i]
		}
		if nums[i] == 0 || nums[i+1] == 0{
			i++
			continue
		}
		
		left, right := i, i+1
		currProduct := nums[left]
		for right<len(nums) && nums[right] != 0 {
			currProduct = currProduct*nums[right]
			if currProduct > maxProduct {
				maxProduct = currProduct
			}
			right++
		}

		for left<right {
			currProduct = currProduct/nums[left]
			if currProduct > maxProduct {
				maxProduct = currProduct
			}
			left++
		}

		i = right+1
	}
	if nums[len(nums)-1] > maxProduct {
		maxProduct = nums[len(nums)-1]
	}

	return maxProduct
}
