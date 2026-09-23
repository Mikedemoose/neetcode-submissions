func myPow(x float64, n int) float64 {
    var res float64
	if x == 0 {
		return res
	} else if n == 0 || x == 1 {
		return 1
	} else if x == -1 {
		if n %2 == 0 {
			return 1
		}
		return -1
	}
	if n > 0 {
		res = x
		for range n-1 {
			res *= x
		}
	} else if n < 0 {
		n = -n
		res = 1/x
		for range n-1 {
			if res == 0 {
				return 0
			}
			res /= x
		}
	}

	return res
}
