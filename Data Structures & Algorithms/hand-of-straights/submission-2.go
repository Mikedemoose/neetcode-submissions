func isNStraightHand(hand []int, groupSize int) bool {
    sort.Slice(hand, func(i, j int)bool{
		return hand[i] < hand[j]
	})

	queue := make([]*QueueItem, 0)

	queueStart := 0
	i := 0
	for i < len(hand) {
		createNew := true
		for j := queueStart; j < len(queue); j++ {
			group := queue[j]
			if hand[i] == group.lastVal + 1 {
				group.lastVal++
				group.size++
				if group.size == groupSize {
					queueStart++
				}
				i++
				createNew = false
			} else if hand[i] > group.lastVal + 1 {
				return false
			}
		}

		if i < len(hand) && createNew {
			if groupSize > 1 {
				queue = append(queue, &QueueItem {
					lastVal: hand[i],
					size: 1,
				})
			}
			i++
		}
	}

	return len(queue) == queueStart

}

type QueueItem struct {
	lastVal int
	size int
}
