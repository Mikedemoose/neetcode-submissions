func isValidOpenClose(o, c rune) bool {
	if (o == '(' && c == ')') || (o == '[' && c == ']') || (o == '{' && c == '}') {
		return true
	}
	return false
}

func isOpen(r rune) bool {
	if r == '(' || r == '[' || r == '{' {
		return true
	}

	return false
}

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
    stack := make([]rune, 0)

	if !isOpen(rune(s[0])) {
		return false
	}

	for _, r := range s {
		if isOpen(r) {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !isValidOpenClose(curr, r) {
				return false
			}
		}
	}

	return len(stack) == 0
}
