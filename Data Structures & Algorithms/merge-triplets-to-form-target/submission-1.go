func mergeTriplets(triplets [][]int, target []int) bool {
    x, y, z := -1, -1, -1
	for _, item := range triplets {
		allowMerge := (item[0] <= target[0]) && (item[1] <= target[1]) && (item[2] <= target[2])
		if allowMerge {
			x = max(x, item[0])
			y = max(y, item[1])
			z = max(z, item[2])
		}

		if (x==target[0]) && (y==target[1]) && (z==target[2]) {
			return true
		}
	}

	return false
}
