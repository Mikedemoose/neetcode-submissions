func findKthLargest(nums []int, k int) int {
	minHeap := MinHeap {
		maxSize: k,
	}

	for _, num := range nums {
		minHeap.CheckAndInsert(num)
	}

	return minHeap.heap[0]
}


func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

type MinHeap struct {
	heap []int
	maxSize int
}

func (m *MinHeap) CheckAndInsert(val int) {
	if len(m.heap) == m.maxSize {
		if m.heap[0] < val {
			m.Pop()
		} else {
			return
		}
	}
	m.Insert(val)
}

func (m *MinHeap) Insert(val int) {
	m.heap = append(m.heap, val)

	curr := len(m.heap)-1
	par := getParentIndex(curr)

	for par >= 0 {
		if m.heap[curr] < m.heap[par] {
			m.heap[curr], m.heap[par] = m.heap[par], m.heap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}
}

func (m *MinHeap) Pop () {
	m.heap[0] = m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]

	m.heapify(0)
}

func (m *MinHeap) heapify(index int) {
	for {
		minIndex := index
		left, right := 2*index+1, 2*index+2

		if left < len(m.heap) && m.heap[left] < m.heap[minIndex] {
			minIndex = left
		}
		if right < len(m.heap) && m.heap[right] < m.heap[minIndex] {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		m.heap[index], m.heap[minIndex] = m.heap[minIndex], m.heap[index]
		index = minIndex
	}
} 