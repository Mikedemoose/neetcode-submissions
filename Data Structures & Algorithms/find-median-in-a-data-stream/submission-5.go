type MedianFinder struct {
    minHeap []int
	minHeapSize int

	maxHeap []int
	maxHeapSize int
}


func Constructor() MedianFinder {
    return MedianFinder{
		minHeap: make([]int, 0),
		maxHeap: make([]int, 0),
	}
}

func (this MedianFinder) printHeaps() {
	fmt.Println("maxHeap:", this.maxHeap)
	fmt.Println("minHeap:", this.minHeap)
	fmt.Println(".................")
}

func (this *MedianFinder) AddNum(num int)  {
	if this.minHeapSize == 0 && this.maxHeapSize == 0 {
		this.minHeapInsert(num)
		// this.printHeaps()
		return
	} 
	var minHeapRoot, maxHeapRoot int
	if this.minHeapSize > 0 {
		minHeapRoot = this.Peek(true)
	} else {
		maxHeapRoot = this.Peek(false)
		if num < maxHeapRoot {
			this.ShiftToMinHeap()
			this.maxHeapInsert(num)
		} else {
			this.minHeapInsert(num)
		}
		// this.printHeaps()
		return
	}
	if this.maxHeapSize > 0 {
		maxHeapRoot = this.Peek(false)
	} else {
		minHeapRoot = this.Peek(true)
		if num > minHeapRoot {
			this.ShiftToMaxHeap()
			this.minHeapInsert(num)
		} else {
			this.maxHeapInsert(num)
		}
		// this.printHeaps()
		return
	}

	if this.minHeapSize > this.maxHeapSize {
		if num > minHeapRoot {
			this.ShiftToMaxHeap()
			this.minHeapInsert(num)
		} else {
			this.maxHeapInsert(num)
		}
	} else {
		if num < maxHeapRoot {
			this.ShiftToMinHeap()
			this.maxHeapInsert(num)
		} else {
			this.minHeapInsert(num)
		}
	}
	// this.printHeaps()
}


func (this *MedianFinder) FindMedian() float64 {
    if this.minHeapSize == this.maxHeapSize {
		return float64(this.Peek(true))/2 + float64(this.Peek(false))/2
	}
	if this.minHeapSize > this.maxHeapSize {
		return float64(this.Peek(true))
	}
	return float64(this.Peek(false))
}

func (this *MedianFinder) ShiftToMinHeap() {
	this.minHeapInsert(this.maxHeapExtractRoot())
}
func (this *MedianFinder) ShiftToMaxHeap() {
	this.maxHeapInsert(this.minHeapExtractRoot())
}


func (this *MedianFinder) minHeapInsert(val int) {
	this.minHeap = append(this.minHeap, val)
	this.minHeapSize++

	curr_index := this.minHeapSize-1
	par_index := getParentIndex(curr_index)

	for curr_index > 0 && this.minHeap[curr_index] < this.minHeap[par_index] {
		this.minHeap[curr_index], this.minHeap[par_index] = this.minHeap[par_index], this.minHeap[curr_index]
		curr_index = par_index
		par_index = getParentIndex(curr_index)
	}
}

func (this *MedianFinder) maxHeapInsert(val int) {
	this.maxHeap = append(this.maxHeap, val)
	this.maxHeapSize++

	curr_index := this.maxHeapSize-1
	par_index := getParentIndex(curr_index)

	for curr_index > 0 && this.maxHeap[curr_index] > this.maxHeap[par_index] {
		this.maxHeap[curr_index], this.maxHeap[par_index] = this.maxHeap[par_index], this.maxHeap[curr_index]
		curr_index = par_index
		par_index = getParentIndex(curr_index)
	}
}

func (this *MedianFinder) minHeapExtractRoot() int {

	if this.minHeapSize == 0 {
		return -1
	}

	head := this.minHeap[0]
	this.minHeap[0] = this.minHeap[this.minHeapSize-1]
	this.minHeapSize--
	this.minHeapify(0)
	this.minHeap = this.minHeap[:this.minHeapSize]

	return head
}

func (this *MedianFinder) maxHeapExtractRoot() int {

	if this.maxHeapSize == 0 {
		return -1
	}

	head := this.maxHeap[0]
	this.maxHeap[0] = this.maxHeap[this.maxHeapSize-1]
	this.maxHeapSize--
	this.maxHeapify(0)
	this.maxHeap = this.maxHeap[:this.maxHeapSize]
	
	return head
}

func (this *MedianFinder) maxHeapify(index int) {
	for {
		maxHeapIndex := index
		left, right := 2*index+1, 2*index+2

		if left < this.maxHeapSize && this.maxHeap[left] > this.maxHeap[maxHeapIndex] {
			maxHeapIndex = left
		}
		if right < this.maxHeapSize && this.maxHeap[right] > this.maxHeap[maxHeapIndex] {
			maxHeapIndex = right
		}
		if index == maxHeapIndex {
			break
		}
		this.maxHeap[maxHeapIndex], this.maxHeap[index] = this.maxHeap[index], this.maxHeap[maxHeapIndex]
		index = maxHeapIndex
	}
}

func (this *MedianFinder) minHeapify(index int) {
	for {
		minHeapIndex := index
		left, right := 2*index+1, 2*index+2

		if left < this.minHeapSize && this.minHeap[left] < this.minHeap[minHeapIndex] {
			minHeapIndex = left
		}
		if right < this.minHeapSize && this.minHeap[right] < this.minHeap[minHeapIndex] {
			minHeapIndex = right
		}
		if index == minHeapIndex {
			break
		}
		this.minHeap[minHeapIndex], this.minHeap[index] = this.minHeap[index], this.minHeap[minHeapIndex]
		index = minHeapIndex
	}
}

func (this MedianFinder) Peek(isMin bool) int {
	if isMin && this.minHeapSize > 0{
		return this.minHeap[0]
	} else if !isMin && this.maxHeapSize > 0 {
		return this.maxHeap[0]
	}
	return -1
}

func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}



