/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := getMinHeap(lists)

	var head, curr *ListNode
	
	// merge the lists 

	for minHeap.Size() > 0 {

		item := minHeap.Pop()

		if head == nil {
			head = item
			curr = head
		} else {
			curr.Next = item
			curr = curr.Next
		}
	}
	if curr != nil {
		curr.Next = nil
	}

	return head
}

type MinHeap struct {
	heap []*ListNode
}

func getParentIndex(c int) int {
	return (c-1)/2
}

func (this MinHeap) Size() int {
	return len(this.heap)
}

func (this *MinHeap) Insert(val *ListNode) {
	this.heap = append(this.heap, val)
	curr := len(this.heap)-1
	par := getParentIndex(curr)

	for curr > par {
		if this.heap[curr].Val < this.heap[par].Val {
			this.heap[curr], this.heap[par] = this.heap[par], this.heap[curr]
			curr = par
			par = getParentIndex(curr)
		} else {
			break
		}
	}
}

func (this *MinHeap) Pop() *ListNode {
	head := this.heap[0]

	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]

	this.heapify(0)

	return head
}

func (this *MinHeap) heapify(index int) {
	for {
		minIndex := index
		left, right := 2*index+1, 2*index+2

		if left < len(this.heap) && this.heap[left].Val < this.heap[minIndex].Val {
			minIndex = left
		}
		if right < len(this.heap) && this.heap[right].Val < this.heap[minIndex].Val {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		this.heap[minIndex], this.heap[index] = this.heap[index], this.heap[minIndex]
		index = minIndex
	}
}

func getMinHeap(lists []*ListNode) MinHeap {
	val := MinHeap {
		heap: make([]*ListNode, 0),
	}

	for _, list := range lists {
		curr := list
		for curr != nil {
			val.Insert(curr)
			curr = curr.Next
		}
	}

	return val
}