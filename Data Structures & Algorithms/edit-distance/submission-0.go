func minDistance(word1 string, word2 string) int {

	word, target := word1, word2
	memo := make(map[MemoKey]int)

	var recurse func(int, int) int

	recurse = func (start1, start2 int) int {
		if start1 == len(word) || start2 == len(target) {
			return max(len(word)-start1, len(target)-start2)
		}
		memoKey := MemoKey {left: start1, right: start2}

		if val, ok := memo[memoKey]; ok {
			return val
		}

		if word[start1] == target[start2] {
			return recurse(start1+1, start2+1)
		}

		// insert a character
		val1 := recurse(start1, start2+1) + 1
		// remove a character
		val2 := recurse(start1+1, start2) + 1
		// replace a character
		val3 := recurse(start1+1, start2+1) + 1

		res := min(val1, val2, val3)
		memo[memoKey] = res
		
		return res
	}


    return recurse(0, 0)
}

type MemoKey struct {
	left int
	right int
}
