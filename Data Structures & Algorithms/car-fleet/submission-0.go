func carFleet(target int, position []int, speed []int) int {
	vals := make([][2]int, 0)
	for i := range len(position) {
		vals = append(vals, [2]int{position[i], speed[i]})
	}

	sort.Slice(vals, func(i, j int)bool {
		return vals[i][0] > vals[j][0]
	})

	tRemaining := float64(target-vals[0][0])/float64(vals[0][1])
	numFleets := 1

	for i:=1; i<len(vals); i++ {
		currTRemaining := float64(target-vals[i][0])/float64(vals[i][1])
		if currTRemaining > tRemaining {
			tRemaining = currTRemaining
			numFleets++
		}
	}

	return numFleets
}
