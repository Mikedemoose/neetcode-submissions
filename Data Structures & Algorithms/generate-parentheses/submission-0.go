func generateParenthesis(n int) []string {
	return recurse(n, n)
}

func recurse(openRem, closeRem int) []string {
	// closeRem >= openRem
	res := make([]string, 0)
	if closeRem == 0 || openRem < 0 {
		return res
	}

	if closeRem >= openRem {
		if closeRem > openRem {
			subRes1 := recurse(openRem, closeRem-1)
			for _, item := range subRes1 {
				res = append(res, ")" + item)
			}
			if len(subRes1) == 0 {
				res = append(res, ")")
			}
		}
		if openRem > 0 {
			subRes2 := recurse(openRem-1, closeRem)
			for _, item := range subRes2 {
				res = append(res, "(" + item)
			}
		}
	}

	return res
}
