type KthLargest struct {
    minHeap []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
    val := KthLargest {
		minHeap : make([]int, 0),
		k: k,
	}

	for _, num := range nums {
		val.Insert(num)
	}

	return val
}


func (this *KthLargest) Add(val int) int {
    this.Insert(val)
	return this.minHeap[0]
}

func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

func (this *KthLargest) Insert(val int) {
	if len(this.minHeap) == this.k {
		if val < this.minHeap[0] {
			return
		}
	}
	this.minHeap = append(this.minHeap, val)
	curr := len(this.minHeap)-1
	par := getParentIndex(curr)

	for curr > par {
		if this.minHeap[curr] < this.minHeap[par] {
			this.minHeap[curr], this.minHeap[par] = this.minHeap[par], this.minHeap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}

	if len(this.minHeap) > this.k {
		this.minHeap[0] = this.minHeap[this.k]
		this.minHeap = this.minHeap[:this.k]

		this.heapify(0)
	}
}

func (this *KthLargest) heapify(index int) {
	for {
		minIndex := index
		left, right := 2*index + 1, 2*index + 2

		if left < len(this.minHeap) && this.minHeap[left] < this.minHeap[minIndex] {
			minIndex = left
		}
		if right < len(this.minHeap) && this.minHeap[right] < this.minHeap[minIndex] {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		this.minHeap[minIndex], this.minHeap[index] = this.minHeap[index], this.minHeap[minIndex]
		index = minIndex
	}
}
