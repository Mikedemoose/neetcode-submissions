func plusOne(digits []int) []int {
    res := make([]int, 0)

	carry := 1

	for i:=len(digits)-1; i>=0; i-- {
		val := digits[i] + carry
		carry = val/10
		res = append([]int{val%10}, res...)
	}

	if carry == 1 {
		res = append([]int{1}, res...)
	}

	return res
}
