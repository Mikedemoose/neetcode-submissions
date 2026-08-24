type PrefixTree struct {
	letters map[rune]*PrefixTree
	isEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree {
		letters: make(map[rune]*PrefixTree),
	}
}

func (this *PrefixTree) Insert(word string) {
	temp := this
	for _, r := range word {
		if val, ok := temp.letters[r]; ok {
			temp = val
		} else {
			newTree := Constructor()
			temp.letters[r] = &newTree
			temp = temp.letters[r]
		}
	}
	temp.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	temp := this
	for _, r := range word {
		if _, ok := temp.letters[r]; !ok {
			return false
		}
		temp = temp.letters[r]
	}
	return temp.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	temp := this
	for _, r := range prefix {
		if _, ok := temp.letters[r]; !ok {
			return false
		}
		temp = temp.letters[r]
	}
	return true
}
