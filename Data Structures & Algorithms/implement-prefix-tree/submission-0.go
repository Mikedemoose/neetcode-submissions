type PrefixTree struct {
	letters map[rune]*PrefixTree
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
	if _, ok := temp.letters['1']; !ok {
		endTree := Constructor()
		temp.letters['1'] = &endTree
	}
}

func (this *PrefixTree) Search(word string) bool {
	temp := this
	for _, r := range word {
		if _, ok := temp.letters[r]; !ok {
			return false
		}
		temp = temp.letters[r]
	}
	if _, ok := temp.letters['1']; ok {
		return true
	}
	return false
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
