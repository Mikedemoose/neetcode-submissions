/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	sort.Slice(intervals, func(i, j int)bool {
		return intervals[i].start < intervals[j].start
	})

	minHeap := MinHeap{}
	maxSize := 0

	for _, interval := range intervals {

		if minInterval, heapSize := minHeap.Peek(); heapSize == 0 {
			minHeap.Insert(interval.end)
		} else {
			if interval.start < minInterval {
				minHeap.Insert(interval.end)
			} else {
				for {
					val, size := minHeap.Peek()
					if size == 0 || val > interval.start {
						break
					}
					_ = minHeap.ExtractMin()
				}
				minHeap.Insert(interval.end)
			}
		}
		if minHeap.size > maxSize {
			maxSize = minHeap.size
		}

	}
	

	return maxSize
}

type MinHeap struct {
	heap []int
	size int
}

func findParentIndex(child int) int {
	return (child-1)/2
}

func (h *MinHeap) Insert(val int) {
	h.heap = append(h.heap, val)
	h.size++

	curr_index := h.size-1
	par_index := findParentIndex(curr_index)

	for curr_index > par_index {
		if h.heap[curr_index] < h.heap[par_index] {
			h.heap[curr_index], h.heap[par_index] = h.heap[par_index], h.heap[curr_index]
			curr_index = par_index
			par_index = findParentIndex(curr_index)
		} else {
			break
		}
	}
}

func (h MinHeap) Peek() (int, int) {
	if h.size > 0 {
		return h.heap[0], h.size
	}
	return -1, 0
}

func (h *MinHeap) ExtractMin() int {

	if h.size == 0 {
		return -1
	}

	minVal := h.heap[0]

	h.heap[0] = h.heap[h.size-1]
	h.size--
	h.siftDown(0)
	h.heap = h.heap[:h.size]

	return minVal
}


func (h *MinHeap) siftDown(index int) {
	for {
		minIndex, left, right := index, 2*index+1, 2*index+2

		if left < h.size && h.heap[left] < h.heap[minIndex] {
			minIndex = left
		}
		if right < h.size && h.heap[right] < h.heap[minIndex] {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		h.heap[minIndex], h.heap[index] = h.heap[index], h.heap[minIndex]
		index = minIndex
	}
}
