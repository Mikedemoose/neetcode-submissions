func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	memo := make(map[MemoKey]int)

    return recurse(amount, coins, 0, memo)
}

func recurse(target int, coins []int, start int, memo map[MemoKey]int) int {
	if start >= len(coins) || target < 0 {
		return 0
	}
	memoKey := MemoKey {
		start: start,
		target: target,
	}
	if val, ok := memo[memoKey]; ok {
		return val
	}

	if target == coins[start] {
		memo[memoKey] = 1
		return 1
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
