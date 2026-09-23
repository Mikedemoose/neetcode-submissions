func networkDelayTime(times [][]int, n int, k int) int {
    queue := NewHeap()

	nodeEdgeMap := make(map[int][][]int)
	for _, time := range times {
		nodeEdgeMap[time[0]] = append(nodeEdgeMap[time[0]], time)
	}

	totalTime := 0
	currTime := 0
	nodeVisitedMap := map[int]struct{} {
		k: struct{}{},
	}

	if val, ok := nodeEdgeMap[k]; ok {
		for _, item := range val {
			queue.Insert(item)
		}
	} else {
		return -1
	}

	for queue.Size() > 0 {
		currItem := queue.Pop()
		if _, ok := nodeVisitedMap[currItem[1]]; !ok {
			totalTime += (currItem[2]-currTime)
			nodeVisitedMap[currItem[1]] = struct{}{}
			currTime = currItem[2]

			if val, ok1 := nodeEdgeMap[currItem[1]]; ok1 {
				for _, item := range val {
					item[2] += currTime
					queue.Insert(item)
				}
			}
		}
	}

	if len(nodeVisitedMap) == n {
		return totalTime
	}
	
	return -1

}

type MinHeap struct {
	heap [][]int
}

func NewHeap() *MinHeap {
	return &MinHeap {
		heap: make([][]int, 0),
	}
}

func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

func (h *MinHeap) Insert(val []int) {
	h.heap = append(h.heap, val)
	curr := len(h.heap)-1
	par := getParentIndex(curr)

	for curr > par {
		if h.heap[curr][2] < h.heap[par][2] {
			h.heap[curr], h.heap[par] = h.heap[par], h.heap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}
}

func (h *MinHeap) Pop() []int {
	head := h.heap[0]

	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	h.heapify(0)

	return head
}

func (h *MinHeap) heapify(index int) {
	for {
		minIndex := index
		left, right := 2*index+1, 2*index+2

		if left < len(h.heap) && h.heap[left][2] < h.heap[minIndex][2] {
			minIndex = left
		}
		if right < len(h.heap) && h.heap[right][2] < h.heap[minIndex][2] {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		h.heap[minIndex], h.heap[index] = h.heap[index], h.heap[minIndex]
		index = minIndex
	}
}

func (h MinHeap) Size() int {
	return len(h.heap)
}
