func minCostConnectPoints(points [][]int) int {
	visited := make(map[Point]struct{})

	totalCost := 0

	currPoint := getPoint(points[0])
	visited[currPoint] = struct{}{}

	minHeap := NewHeap()

	for i:=1; i<len(points); i++ {
		minHeap.Insert(HeapItem{targetNode: getPoint(points[i]), distance:getDistanceBetweenPoints(points[0], points[i])})
	}

	for minHeap.Size() > 0 {
		head := minHeap.Pop()

		if _, ok := visited[head.targetNode]; ok {
			continue
		}

		totalCost += head.distance
		visited[head.targetNode] = struct{}{}
		p := []int{head.targetNode.x, head.targetNode.y}

		for _, point := range points {
			currPoint := getPoint(point)
			if _, ok := visited[currPoint]; ok {
				continue
			}
			minHeap.Insert(HeapItem{targetNode: getPoint(point), distance: getDistanceBetweenPoints(p, point)})
		}
	}


	return totalCost
}

func getDistanceBetweenPoints(pointA, pointB []int) int {
	x_dist := pointA[0] - pointB[0]
	y_dist := pointA[1] - pointB[1]

	if x_dist < 0 {
		x_dist = -x_dist
	}
	if y_dist < 0 {
		y_dist = -y_dist
	}

	return x_dist + y_dist
}

type MinHeap struct {
	heap []HeapItem
}

type HeapItem struct {
	targetNode Point
	distance int
}

type Point struct {
	x int
	y int
}

func getPoint(val []int) Point {
	return Point {
		x: val[0],
		y: val[1],
	}
}

func getParentIndex(childIndex int) int {
	return (childIndex-1)/2
}

func (this *MinHeap) Insert(val HeapItem) {
	this.heap = append(this.heap, val)
	curr := len(this.heap)-1
	par := getParentIndex(curr)

	for curr > par {
		if this.heap[curr].distance < this.heap[par].distance {
			this.heap[curr], this.heap[par] = this.heap[par], this.heap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}
}

func (this *MinHeap) Pop() HeapItem {
	head := this.heap[0]

	currLen := len(this.heap)
	this.heap[0] = this.heap[currLen-1]
	this.heap = this.heap[:currLen-1]

	this.heapify(0)

	return head
}

func (this *MinHeap) heapify(index int) {
	for {
		minIndex := index
		left, right := 2*index+1, 2*index+2

		if left < len(this.heap) && this.heap[left].distance < this.heap[minIndex].distance {
			minIndex = left
		}
		if right < len(this.heap) && this.heap[right].distance < this.heap[minIndex].distance {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		this.heap[minIndex], this.heap[index] = this.heap[index], this.heap[minIndex]
		index = minIndex
	}
}

func (this MinHeap) Size() int {
	return len(this.heap)
}

func NewHeap() *MinHeap {
	return &MinHeap {
		heap: make([]HeapItem, 0),
	}
}
