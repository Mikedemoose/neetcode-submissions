func countOnes(n int) int {
	total := 0
	for n > 0 {
		total += n%2
		n >>=1
	}

	return total
}


func countBits(n int) []int {
	result := make([]int, n+1)

	for i:=1; i<n+1; i++ {
		result[i] = countOnes(i)
	}

	return result
}
