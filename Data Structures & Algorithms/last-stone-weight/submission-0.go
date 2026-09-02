func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}

	for _, stone := range stones {
		h.Insert(stone)
	}

	for h.size > 1 {
		p1 := h.Pop()
		p2 := h.Pop()

		if p1 == p2 {
			continue
		}

		if p1 > p2 {
			h.Insert(p1-p2)
		} else {
			h.Insert(p2-p1)
		}
	}

	if h.size == 1 {
		return h.Pop()
	}

	return 0
}

type MaxHeap struct {
	maxHeap []int
	size int
}

func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

func (h *MaxHeap) Insert(val int) {
	h.maxHeap = append(h.maxHeap, val)
	h.size++

	curr_index, par_index := h.size-1, getParentIndex(h.size-1)
	for par_index >= 0 {
		if h.maxHeap[par_index] >= h.maxHeap[curr_index] {
			break
		}
		h.maxHeap[par_index], h.maxHeap[curr_index] = h.maxHeap[curr_index], h.maxHeap[par_index]
		curr_index = par_index
		par_index = getParentIndex(curr_index)
	}
}

func (h *MaxHeap) Pop() int {

	if h.size == 0 {
		return 0
	}

	head := h.maxHeap[0]
	h.maxHeap[0] = h.maxHeap[h.size-1]
	h.size--
	h.maxHeap = h.maxHeap[:h.size]

	h.heapify(0)
	return head
}


func (h *MaxHeap) heapify(index int) {
	for {
		max_index := index
		left, right := 2*index+1, 2*index+2

		if left < h.size {
			if h.maxHeap[left] > h.maxHeap[max_index] {
				max_index = left
			}
		}
		if right < h.size {
			if h.maxHeap[right] > h.maxHeap[max_index] {
				max_index = right
			}
		}

		if max_index == index {
			break
		}

		h.maxHeap[max_index], h.maxHeap[index] = h.maxHeap[index], h.maxHeap[max_index]
		index = max_index
	}
}
