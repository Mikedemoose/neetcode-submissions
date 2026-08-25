func lengthOfLIS(nums []int) int {
    lisArray := make([]int, len(nums))

	lisArray[0] = 1

	maxVal := 1

	for i:=1; i<len(nums); i++ {
		lisArray[i] = 1

		for j:=i-1; j>=0; j-- {
			if (nums[j] < nums[i]) && (lisArray[j] >= lisArray[i]) {
				lisArray[i] = lisArray[j] + 1
			}

			if lisArray[i] > maxVal {
				maxVal = lisArray[i]
			}
		}
	}

	return maxVal
}
