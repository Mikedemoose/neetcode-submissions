type MedianFinder struct {
    nums []int
}


func Constructor() MedianFinder {
    return MedianFinder{
		nums: make([]int, 0),
	}
}


func (this *MedianFinder) AddNum(num int)  {
    this.nums = append(this.nums, num)

	sort.Slice(this.nums, func(i, j int) bool {
		return this.nums[i] < this.nums[j]
	})
}


func (this *MedianFinder) FindMedian() float64 {
    size := len(this.nums)

	mid := (size+1)/2
	if size & 1 == 1 {
		return float64(this.nums[mid-1])
	}
	return float64(this.nums[mid-1]+this.nums[mid])/2
}





// type MaxHeap {
// 	heap []int
// 	size int
// 	sortedArray []int
// }

// func getParentIndex(childIndex int) int {
// 	return (childIndex-1)/2
// }

// func (h *MaxHeap) Insert(item int) {
// 	h.heap = append(h.heap, item)
// 	h.sortedArray = append(h.sortedArray, item)
// 	h.size++

// 	curr_index = h.size-1
// 	par_index = getParentIndex(curr_index)

// 	for curr_index > par_index {
// 		if h.heap[curr_index] > h.heap[par_index] {
// 			h.heap[curr_index], h.heap[par_index] = h.heap[par_index], h.heap[curr_index]
// 		} else {
// 			break
// 		}
// 	}

// 	h.ComputeSortedArray()
// }

// func (h *MaxHeap) ComputeSortedArray() {

// }
