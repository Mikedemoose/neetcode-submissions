func change(amount int, coins []int) int {
	memo := make(map[MemoKey]int)
    return recurse(amount, coins, 0, memo)
}

func recurse(target int, coins []int, start int, memo map[MemoKey]int) int {
	if start >= len(coins) || target < 0 {
		return 0
	} else if target == 0 {
		return 1
	}
	
	memoKey := MemoKey {
		start: start,
		target: target,
	}
	if val, ok := memo[memoKey]; ok {
		return val
	}

	withStart := recurse(target-coins[start], coins, start, memo)
	withoutStart := recurse(target, coins, start+1, memo)

	total := withStart + withoutStart
	memo[memoKey] = total

	return total
}

type MemoKey struct {
	start int
	target int
}
