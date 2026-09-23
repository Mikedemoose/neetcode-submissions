func isNStraightHand(hand []int, groupSize int) bool {
    sort.Ints(hand)
	counts := make(map[int]int)
	for _, item := range hand {
		counts[item]++
	}


	for _, item := range hand {
		if counts[item] == 0 {
			continue
		}

		val := item
		for range groupSize {
			if c, ok := counts[val]; ok {
				if c <= 0 {
					return false
				}
				counts[val]--
				val++
			} else {
				return false
			}
		}
	}


	return true
}
