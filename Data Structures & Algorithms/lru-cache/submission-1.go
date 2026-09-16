type LRUCache struct {
    
	cacheHead *Node
	cacheTail *Node
	currLength int
	maxLength int

	keyMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache {
		maxLength: capacity,
		keyMap: make(map[int]*Node),
	}
    
}

func (this *LRUCache) Get(key int) int {
	val := -1
    if node, ok := this.keyMap[key]; ok {
		val = node.val
		this.removeKey(key)
		this.appendKey(key, val)
	}
	return val
}

func (this *LRUCache) Put(key int, value int) {
	this.removeKey(key)
	this.appendKey(key, value)

	if this.currLength > this.maxLength {
		this.removeKey(this.cacheHead.key)
	}
}

func (this *LRUCache) removeKey(key int) {
	if node, ok := this.keyMap[key]; ok {
		// if key is terminal
		if node == this.cacheHead || node == this.cacheTail {
			if node == this.cacheHead {
				this.cacheHead = node.next
				if this.cacheHead != nil {
					this.cacheHead.prev = nil
				}
			}
			if node == this.cacheTail {
				this.cacheTail = node.prev
				if this.cacheTail != nil {
					this.cacheTail.next = nil
				}
			}
		} else {
		// if key is non-terminal
			node.prev.next = node.next
			node.next.prev = node.prev
		}
		delete(this.keyMap, key)
		this.currLength--
	}
}

func (this *LRUCache) appendKey(key int, value int) {

	newNode := &Node {
		key: key,
		val: value,
	}

	if this.currLength == 0 {
		this.cacheHead, this.cacheTail = newNode, newNode
	} else {
		newNode.prev = this.cacheTail
		this.cacheTail.next = newNode
		this.cacheTail = newNode
	}


	this.currLength++
	this.keyMap[key] = newNode
}

type Node struct {
	key int
	val int
	prev *Node
	next *Node
}
