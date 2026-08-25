func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left <= right {
		l, r := s[left], s[right]
		if l>='A' && l <= 'Z' {
			l = l-'A'+'a'
		} else if (l>='a' && l<='z') || (l>='0' && l<='9') {
			l = l
		} else {
			left++
			continue
		}

		if r>='A' && r <= 'Z' {
			r = r-'A'+'a'
		} else if (r>='a' && r<='z') || (r>='0' && r<='9') {
			r = r
		} else {
			right--
			continue
		}

		if l != r {
			return false
		}
		left++
		right--
	}

	return true
}
