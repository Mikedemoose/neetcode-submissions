func plusOne(digits []int) []int {
    res := append([]int{}, digits...)

	carry := 1
	index := len(digits)-1

	for carry == 1 && index >= 0 {
		val := res[index] + carry
		carry = val/10
		res[index] = val%10

		index--
	}

	if carry == 1 {
		res = append([]int{1}, res...)
	}

	return res
}
