func lengthOfLIS(nums []int) int {

	longest := 1

	subSeqLength := make([]int, len(nums))
	subSeqLength[0] = 1

	for i := 1; i<len(nums); i++ {
		currLongest := 1

		for j := 0; j<i; j++ {
			if nums[j] < nums[i] && subSeqLength[j]+1 > currLongest {
				currLongest = subSeqLength[j]+1
			}
		}

		subSeqLength[i] = currLongest
		if currLongest > longest {
			longest = currLongest
		}

	}

	return longest
    
}
