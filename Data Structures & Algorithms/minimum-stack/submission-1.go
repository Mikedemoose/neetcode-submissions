type MinStack struct {
	stack []int
	mins []int
	currentMin int
}

func Constructor() MinStack {
	return MinStack {
		stack: make([]int, 0),
		mins: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.stack)==1 || val < this.currentMin {
		this.currentMin = val
	}
	this.mins = append(this.mins, this.currentMin)
}

func (this *MinStack) Pop() {
	l := len(this.stack)
	this.stack = this.stack[:l-1]
	this.mins = this.mins[:l-1]
	l--
	if l>0 && this.mins[l-1] > this.currentMin {
		this.currentMin = this.mins[l-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.currentMin
}
