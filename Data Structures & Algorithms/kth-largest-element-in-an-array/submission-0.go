func findKthLargest(nums []int, k int) int {
	heapSize := k

	maxHeap := MaxHeap {
		maxSize: heapSize,
	}

	for _, num := range nums {
		maxHeap.Insert(num)
	}

	var val int
	for range k {
		val = maxHeap.GetTop()
	}

	return val
}


func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

type MaxHeap struct {
	heap []int
	maxSize int
}

func (m *MaxHeap) Insert(val int) {
	m.heap = append(m.heap, val)

	curr := len(m.heap)-1
	par := getParentIndex(curr)

	for par >= 0 {
		if m.heap[curr] > m.heap[par] {
			m.heap[curr], m.heap[par] = m.heap[par], m.heap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}
}

func (m *MaxHeap) GetTop() int {
	top := m.heap[0]

	m.heap[0] = m.heap[len(m.heap)-1]
	m.heap = m.heap[:len(m.heap)-1]

	m.heapify(0)


	return top
}

func (m *MaxHeap) heapify(index int) {
	for {
		maxIndex := index
		left, right := 2*index+1, 2*index+2

		if left < len(m.heap) && m.heap[left] > m.heap[maxIndex] {
			maxIndex = left
		}
		if right < len(m.heap) && m.heap[right] > m.heap[maxIndex] {
			maxIndex = right
		}

		if maxIndex == index {
			break
		}

		m.heap[maxIndex], m.heap[index] = m.heap[index], m.heap[maxIndex]

		index = maxIndex
	}
}

// type MinHeap struct {
// 	heap []int
// 	maxSize int
// }

// func (m *MinHeap) Insert(val int) {
// 	m.heap = append(m.heap, val)
// 	curr_size := min(len(m.heap), m.maxSize)

// 	curr := len(m.heap)-1
// 	par := getParentIndex(curr)

// 	for par >= 0 {
// 		if m.heap[curr] < m.heap[par] {
// 			m.heap[curr], m.heap[par] = m.heap[par], m.heap[curr]
// 			curr = par
// 			par = getParentIndex(curr)
// 		} else {
// 			break
// 		}
// 	}

// 	m.heap = m.heap[:curr_size]
// }