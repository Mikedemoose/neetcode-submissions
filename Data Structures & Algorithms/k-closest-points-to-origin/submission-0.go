func kClosest(points [][]int, k int) [][]int {
	minHeap := MinHeap {}
	for _, p := range points {
		minHeap.Insert(p)
	}

	res := make([][]int, 0)
	for range k {
		res = append(res, minHeap.GetTop())
	}

	return res
}

func getDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

type MinHeap struct {
	heap [][]int
}

func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

func (m *MinHeap) Insert(val []int) {
	m.heap = append(m.heap, val)

	curr := len(m.heap)-1
	par := getParentIndex(curr)

	for par >= 0 {
		if getDistance(m.heap[curr]) < getDistance(m.heap[par]) {
			m.heap[curr], m.heap[par] = m.heap[par], m.heap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}
}

func (m *MinHeap) GetTop() []int {
	top := m.heap[0]

	m.heap[0] = m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]

	m.heapify(0)

	return top
}

func (m *MinHeap) heapify(index int) {
	for {
		minIndex := index
		left, right := 2*index+1, 2*index+2

		if left < len(m.heap) && getDistance(m.heap[left]) < getDistance(m.heap[minIndex]) {
			minIndex = left
		}
		if right < len(m.heap) && getDistance(m.heap[right]) < getDistance(m.heap[minIndex]) {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		m.heap[minIndex], m.heap[index] = m.heap[index], m.heap[minIndex]
		m.heapify(minIndex)
	}
}
