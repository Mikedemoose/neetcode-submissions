func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1) + len(s2) {
		return false
	}

	memo := make(map[MemoKey]bool)

	return recurse(s1, s2, s3, 0, 0, 0, memo)
    
}

func recurse(s1, s2, s3 string, left, right, curr int, memo map[MemoKey]bool) bool {

	if left >= len(s1) {
		return s2[right:] == s3[curr:]
	} else if right >= len(s2) {
		return s1[left:] == s3[curr:]
	}

	memoKey := MemoKey {
		left: left,
		right: right,
		curr: curr,
	}

	if val, ok := memo[memoKey]; ok {
		return val
	}

	if s1[left] == s3[curr] {
		if recurse(s1, s2, s3, left+1, right, curr+1, memo) {
			return true
		}
	}
	if s2[right] == s3[curr] {
		if recurse(s1, s2, s3, left, right+1, curr+1, memo) {
			return true
		}
	}

	memo[memoKey] = false

	return false
}

type MemoKey struct {
	left int
	right int
	curr int
}
